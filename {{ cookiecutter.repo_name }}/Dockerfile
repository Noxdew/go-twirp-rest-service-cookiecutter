FROM golang:latest as builder

# Setup dependencies
RUN apt-get update && \
    apt-get -y install protobuf-compiler libprotoc-dev git

RUN go get github.com/twitchtv/twirp/protoc-gen-twirp
RUN go get bitbucket.org/noxdew/twirp-rest-gen
# protoc-gen-go has to be the same version as proto package
RUN git clone https://github.com/golang/protobuf.git /go/src/github.com/golang/protobuf; \
    cd "$(go env GOPATH)"/src/github.com/golang/protobuf; \
    git checkout v1.2.0; \
    go install github.com/golang/protobuf/protoc-gen-go

# Build the service
COPY . /go/src/bitbucket.org/{{ cookiecutter.repo_owner }}/{{ cookiecutter.repo_name }}/
WORKDIR /go/src/bitbucket.org/{{ cookiecutter.repo_owner }}/{{ cookiecutter.repo_name }}/
RUN go generate $(go list ./...)
RUN CGO_ENABLED=0 go build -ldflags "-w -extldflags "-static" -X main.version=`git rev-parse HEAD`" -v -o {{ cookiecutter.service_name }} bitbucket.org/{{ cookiecutter.repo_owner }}/{{ cookiecutter.repo_name }}/service

FROM scratch
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /go/src/bitbucket.org/{{ cookiecutter.repo_owner }}/{{ cookiecutter.repo_name }}/{{ cookiecutter.service_name }} /{{ cookiecutter.service_name }}
COPY --from=builder /go/src/bitbucket.org/{{ cookiecutter.repo_owner }}/{{ cookiecutter.repo_name }}/config /config
EXPOSE 8080
CMD [ "/{{ cookiecutter.service_name }}" ]
