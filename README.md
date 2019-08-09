# message-broker-miscellaneous

<img align="center" width="900" height="447" src="https://github.com/enbis/message-broker-miscellaneous/blob/master/image/enbis_mbm_logo.png">

Like a telephone children's game, this tool connects different application protocols (HTTP request, HTTP/2 protocol buffer) and open source messaging systems (NATS MQTT) and allows transfert data between them. From HTTP request to the MQTT message pubblication, passing through HTTP/2 and NATS. It's also possible to run Client/Server system standalone to test and evaluate the communications individually. **Fully developed in GO.**   

### Generate protobuf

- `make generate/go`

### Makefile

- `start/all` to launch docker NATS & MQTT, HTTP Server, gRPC Server and NATS Subscription
- `request/curl` to launch the HTTP request

## Cobra commands

### HTTP Server

- `go run main.go httpserver` to launch HTTPServer in listening mode, on port setted in `config.yml`. Default value is `3300`. The server receive the request with params `http://localhost:3300/init?topic=<preferred_topic>&payload=<preferred_payload>`, create the protocl buffer and foreward it to the gRPC Server.

### gRPC Server

- `go run main.go grpcserver` to launch gRPCServer in listening mode, on port setted in `config.yml`. Default value is `7777`.

### gRPC Client

- `go run main.go grpcclient` to interact with gRPCServer, sending protobuffer `models.PingMessage`. The protocol bufer contains `topic` and `payload` that will be needed for the Nats messaging system.

### nats subscriber

- first of all `docker-compose up` to launch the Nats service on port 4222

- `go run main.go natssubscriber --topic=<preferred_topic>` to subscribe to the topic preferred. Be careful that the `--topic` match with the topic selected at the http request time, otherwise you won't see the message coming. If no topic is provided it reads the value from the config.yml. 

### nats publisher

- `go run main.go natspublisher --topic=<preferred_topic> --payload=<preferred_payload>` useful to test nats messaging system standalone. After launching Nats service and nats subscriber, launch the publisher to interact with the message broker. 
