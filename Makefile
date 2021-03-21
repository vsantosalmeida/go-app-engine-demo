BINARY_PATH=./bin
BINARY_NAME=$(BINARY_PATH)/go-app-engine-demo.bin
VERSION=1.0.0

build-local:
	@ echo " ---         BUILDING        --- "
	@ go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(BINARY_NAME)
	@ echo " ---      FINISH BUILD       --- "