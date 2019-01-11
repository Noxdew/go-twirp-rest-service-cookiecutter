# {{ cookiecutter.service_name }}

A GoLang service using the Twirp RPC framework and a REST gateway for it. It has been generated using [go-twirp-rest-service-cookiecutter](https://github.com/Noxdew/go-twirp-rest-service-cookiecutter) template.

## Development

1. Run `dep ensure` to install all dependencies and create the `vendor` directory
2. Run `make` to generate the `.twirp.go` and `.rest.go` files from the protobuf service definition, compile the service and run it

## Testing

1. Run `make setup-test` to install the test dependencies
2. Run `make test` to run the tests

## Running in Docker

The Dockerfile defines a multi-stage build, where the first step is installing `protoc` and all required plugins and builds the service. The second creates an image `FROM scratch` to ensure the smallest image size.

The compose file is the simplest definition possible to deploy using `docker-compose` or to a docker swarm.

## Pipeline integration

The pipeline provided works on Bitbucket Pipelines, however it gives a starting point for defining a pipeline in any CI framework.

## How it works

1. The config dir:

The yaml files in there are the ones loaded by the service and configuring all predefined and custom features. The config loaded can be specified using the `ENV` environmental variable (defaults to `prod`). You can check the repositories for the [logging](https://bitbucket.org/noxdew/log) and [metrics](https://bitbucket.org/noxdew/metrics) libraries to see the list of supported config and well as add your own under the `custom` top level property. You can find out more about the config package in the [repo](https://bitbucket.org/noxdew/config).

2. The proto dir:

Define your service(s) in `.proto` files like the example (more details in [Twirp](https://github.com/twitchtv/twirp)). Then define how these services can be accessed using the REST API of your service (more details in [twirp-rest-gen](https://bitbucket.org/noxdew/twirp-rest-gen)). The `gen.go` file defines the commands used to generate the `.go` files from the `.proto` and `.yml` files.

3. The service dir:

- `main.go` - define more healthchecks, register/rename services.
- `impl/impl.go` - an implementation of the HelloWorld example service. Your business logic will go here. Note the `ServiceConfig` is being populated by the values under the `custom` top level property in your config files.
- `impl/client/client.go` - example client for a service.

## Notes

You may notice that `github.com/golang/protobuf` is locked at version `1.2.0` in `Gopkg.toml` as well as this version is being installed in Dockerfile and the pilelines definition. The reason is `github.com/golang/protobuf/proto` checks that it is the same version as `github.com/golang/protobuf/protoc-gen-go`, however `protoc-gen-go` is not direct dependency of the service so it is not installed by `dep ensure`.
