package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	proto "{{ cookiecutter.repo_host }}/{{ cookiecutter.repo_owner }}/{{ cookiecutter.repo_name }}/proto"
	impl "{{ cookiecutter.repo_host }}/{{ cookiecutter.repo_owner }}/{{ cookiecutter.repo_name }}/service/impl"

	configLoader "bitbucket.org/noxdew/config"
	statter "bitbucket.org/noxdew/metrics/statter"
	mux "github.com/gorilla/mux"
	healthcheck "github.com/heptiolabs/healthcheck"
	twirpStatsd "github.com/twitchtv/twirp/hooks/statsd"
)

func main() {
	customConfig := impl.ServiceConfig{}
	config := configLoader.Load(customConfig)
	logger := config.Logging.Create()
	defer logger.Sync()
	config.Metrics.Configure(logger)

	health := healthcheck.NewHandler()
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))
	// Add your own healthchecks

	router := mux.NewRouter()
	router.HandleFunc("/health/ready", health.ReadyEndpoint)
	router.HandleFunc("/health/live", health.LiveEndpoint)

	hook := twirpStatsd.NewStatsdServerHooks(statter.MetricsStatter{})
	service := impl.CreateService(logger, &customConfig)
	twirpServer := proto.NewGreeterServer(service, hook)
	twirpRestServer := proto.NewTwirpRestServer(twirpServer, router)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: twirpRestServer,
	}
	defer server.Shutdown(nil)
	go func() {
		var err error
		if config.TLS.Cert != "" && config.TLS.Key != "" {
			logger.Info("Starting with TLS")
			tlsConfig := &tls.Config{
				MinVersion:               tls.VersionTLS12,
				CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
				PreferServerCipherSuites: true,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				},
			}
			server.TLSConfig = tlsConfig
			server.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)
			err = server.ListenAndServeTLS(config.TLS.Cert, config.TLS.Key)
		} else {
			logger.Info("Starting without TLS")
			err = server.ListenAndServe()
		}
		if err != nil {
			logger.Panic(err.Error())
		}
		logger.Info("Server started")
	}()

	// Handle interrupts
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	logger.Info("Server stopped")
}
