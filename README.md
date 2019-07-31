# message-broker-miscellaneous

### Generate protobuf

- `make generate/go`

## Cobra commands

### gRPC Server

- `go run main.go grpcserver` to launch gRPCServer in listening mode, on port setted in `config.yml`

### gRPC Client

- `go run main.go grpcclient` to interact with gRPCServer, sending protobuffer `models.PingMessage` which contains `topic` and `payload` obtained from the configuration file `config.yml`