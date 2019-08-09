CWDIR := $(shell pwd)
SRC_DIR := $(CWDIR)/models/proto
DST_DIR := $(CWDIR)/models/src
PROJECT_NAME = $(notdir $(PWD))

TOPIC?=$(shell ./function.sh _get_config_value "topic" )

generate/go:
	mkdir -p $(DST_DIR)
	protoc --proto_path=$(SRC_DIR) \
		--go_out=plugins=grpc:$(DST_DIR) \
		$(shell ls $(SRC_DIR)/*.proto)

start/all: docker/up start/httpserver start/grpcserver nats/subscribe

start/httpserver:
	gnome-terminal -- bash -c "go run main.go httpserver; exec bash"

start/grpcserver:
	gnome-terminal -- bash -c "go run main.go grpcserver; exec bash"

nats/subscribe:
	gnome-terminal -- bash -c "go run main.go natssubscriber --topic=$(TOPIC); exec bash"

request/curl:
	curl -X GET 'http://localhost:3300/init?topic=foo&payload=payload'

docker/up:
	docker-compose -p $(PROJECT_NAME) up -d

docker/down:
	docker-compose -p $(PROJECT_NAME) down --remove-orphans

test:
	@echo $(TOPIC)
