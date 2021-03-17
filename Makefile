MAIN_FILE=./cmd/main.go
BINARY_PATH=./bin
BINARY_NAME=$(BINARY_PATH)/account-activity-timeline-consumer.bin
VERSION=1.0.0

build-local:
	@ echo " ---         BUILDING        --- "
	@ go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(BINARY_NAME) $(MAIN_FILE)
	@ echo " ---      FINISH BUILD       --- "