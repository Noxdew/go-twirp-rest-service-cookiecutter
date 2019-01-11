package impl

import (
	"context"

	proto "{{ cookiecutter.repo_host }}/{{ cookiecutter.repo_owner }}/{{ cookiecutter.repo_name }}/proto"
	twirp "github.com/twitchtv/twirp"
	zap "go.uber.org/zap"
)

// ServiceConfig struct containing the custom config for this service
// it will be parsed from the custom top level key in the config.yaml
type ServiceConfig struct {
	WithName  bool   `yaml:"withName" json:"withName"`
	ClientURL string `yaml:"clientUrl" json:"clientUrl"`
}

// Service struct containing the business logic
type Service struct {
	Logger *zap.Logger
	Config ServiceConfig
}

// SayHello example function
func (h *Service) SayHello(ctx context.Context, helloReq *proto.HelloRequest) (*proto.HelloResponse, error) {
	if h.Config.WithName {
		if len(helloReq.Name) <= 0 {
			return nil, twirp.InvalidArgumentError("Name", "No name provided to SayHello")
		}
		return &proto.HelloResponse{
			Message: "Hello, " + helloReq.Name + "!",
		}, nil
	}
	return &proto.HelloResponse{
		Message: "Hello, World!",
	}, nil
}
