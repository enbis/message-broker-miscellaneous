# message-broker-miscellaneous

### Generate protobuf

- `make generate/go`

## Cobra commands

### gRPC Server

- `go run main.go grpcserver` to launch gRPCServer in listening mode, on port setted in `config.yml`. Default value is `7777`.

### gRPC Client

- `go run main.go grpcclient` to interact with gRPCServer, sending protobuffer `models.PingMessage`. The protocol bufer contains `topic` and `payload` that will be needed for the Nats messaging system. 