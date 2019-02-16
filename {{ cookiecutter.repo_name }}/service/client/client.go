package client

import (
	"context"
	"log"
	"net/http"

	proto "{{ cookiecutter.repo_host }}/{{ cookiecutter.repo_owner }}/{{ cookiecutter.repo_name }}/proto"
	impl "{{ cookiecutter.repo_host }}/{{ cookiecutter.repo_owner }}/{{ cookiecutter.repo_name }}/service/impl"
	twirp "github.com/twitchtv/twirp"
	zap "go.uber.org/zap"
)

// ServiceClient abstracts the communication to the service
type ServiceClient struct {
	Logger *zap.Logger
	Config impl.ServiceConfig
}

// NewClient Creates a new service client
func NewClient(config impl.ServiceConfig, logger *zap.Logger) *ServiceClient {
	return &ServiceClient{
		Logger: logger,
		Config: config,
	}
}

// GetHello calls an external service to retrieve the resource
func (sc *ServiceClient) GetHello(name string) (res *proto.HelloResponse, err error) {
	client := proto.NewGreeterJSONClient(sc.Config.ClientURL, &http.Client{})

	for i := 0; i < 5; i++ {
		res, err = client.SayHello(context.Background(), &proto.HelloRequest{Name: name})
		if err != nil {
			if twerr, ok := err.(twirp.Error); ok {
				if twerr.Meta("retryable") != "" {
					// Log the error and go again.
					sc.Logger.Warn("Failed to call client",
						zap.String("client", "Greeter"),
						zap.String("function", "SayHello"),
						zap.String("url", sc.Config.ClientURL),
						zap.Bool("retryable", true),
						zap.Error(err),
					)
					continue
				}
			}
			sc.Logger.Error("Failed to call client",
				zap.String("client", "Greeter"),
				zap.String("function", "SayHello"),
				zap.String("url", sc.Config.ClientURL),
				zap.Bool("retryable", false),
				zap.Error(err),
			)
			return nil, err
		}
		break
	}
	return res, nil
}
