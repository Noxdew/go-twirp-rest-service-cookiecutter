run: build
	ENV=dev ./{{ cookiecutter.service_name }}

build: generate
	go build -o {{ cookiecutter.service_name }} ./service

generate:
	go generate $$(go list ./...)

setup-test: generate
	go get github.com/vektra/mockery/.../

test:
	go test -v ./...
