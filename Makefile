CWDIR := $(shell pwd)
SRC_DIR := $(CWDIR)/models/proto
DST_DIR := $(CWDIR)/models/src

generate/go:
	mkdir -p $(DST_DIR)
	protoc --proto_path=$(SRC_DIR) \
		--go_out=plugins=grpc:$(DST_DIR) \
		$(shell ls $(SRC_DIR)/*.proto)