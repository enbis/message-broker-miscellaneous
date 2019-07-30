CWDIR := $(shell pwd)
SRC_DIR := $(CWDIR)/models/proto
DST_DIR := $(CWDIR)/models

generate/go:
	mkdir -p $(DST_DIR)/src
	protoc -I=$(SRC_DIR) \
		--go_out=$(DST_DIR)/src \
		$(SRC_DIR)/ping_message.proto