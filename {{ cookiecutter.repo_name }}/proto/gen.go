package proto

//go:generate protoc -I . service.proto --twirp_out=. --go_out=.
//go:generate twirp-rest-gen -s service.yml
