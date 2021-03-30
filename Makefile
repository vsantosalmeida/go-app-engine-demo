BINARY_PATH=./bin
BINARY_NAME=$(BINARY_PATH)/go-app-engine-demo.bin
VERSION=1.0.0
PROJECT=${DEVSHELL_PROJECT_ID}


clean:
	@ rm -rf bin/*

build-api:
	@ echo " ---         BUILDING        --- "
	@ go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(BINARY_NAME) api/main.go
	@ echo " ---      FINISH BUILD       --- "

set-project:
	@ sed -i "s/project/${PROJECT}/g" app.yaml

dependencies:
	@ go mod download

start-datastore:
	@ gcloud beta emulators datastore start --project gcp-app-engine --no-store-on-disk