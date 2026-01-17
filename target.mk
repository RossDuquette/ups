.PHONY: build

GO_BUILD=go build -buildvcs=false
BUILD_DIR=./build/bin

clean:
	@rm -rf $(BUILD_DIR)

build: format build_broker build_pub build_sub

format:
	@go fmt ./...

build_broker:
	$(GO_BUILD) -o $(BUILD_DIR)/ups_broker ./cmd/broker

build_pub:
	$(GO_BUILD) -o $(BUILD_DIR)/ups_pub_cli ./cmd/publisher

build_sub:
	$(GO_BUILD) -o $(BUILD_DIR)/ups_sub_cli ./cmd/subscriber

test: format
	go test -v ./...

lint:
	@go vet
