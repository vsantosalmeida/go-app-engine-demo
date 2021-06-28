BINARY_PATH=./bin
BINARY_NAME=$(BINARY_PATH)/go-app-engine-demo.bin
VERSION=1.0.0
PROJECT=${DEVSHELL_PROJECT_ID}


clean:
	@ rm -rf bin/*

build-api:
	@ echo " ---         BUILDING        --- "
	@ go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(BINARY_NAME) cmd/person/main.go
	@ echo " ---      FINISH BUILD       --- "

set-project:
	@ sed -i "s/project/${PROJECT}/g" app.yaml

start-datastore:
	@ gcloud beta emulators datastore start --project gcp-app-engine --no-store-on-disk

compile-protobuf:
	@ protoc --proto_path=protobuf --go_out=protobuf person.proto