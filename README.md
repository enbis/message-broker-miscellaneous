# message-broker-miscellaneous

### Generate protobuf

- `make generate/go`

## Cobra commands

### HTTP Server

- `go run main.go httpserver` to launch HTTPServer in listening mode, on port setted in `config.yml`. Default value is `3300`. The server receive the request with params `http://localhost:3300/init?topic=<preferred_topic>&payload=<preferred_payload>`, create the protocl buffer and foreward it to the gRPC Server.

### gRPC Server

- `go run main.go grpcserver` to launch gRPCServer in listening mode, on port setted in `config.yml`. Default value is `7777`.

### gRPC Client

- `go run main.go grpcclient` to interact with gRPCServer, sending protobuffer `models.PingMessage`. The protocol bufer contains `topic` and `payload` that will be needed for the Nats messaging system. 