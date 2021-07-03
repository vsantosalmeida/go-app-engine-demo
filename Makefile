BINARY_PATH=./bin
PERSON_BINARY_NAME=$(BINARY_PATH)/person-api.bin
JOB_BINARY_NAME=$(BINARY_PATH)/job.bin
CRYPTO_BINARY_NAME=$(BINARY_PATH)/crypto-api.bin
VERSION=1.0.0
PROJECT=${DEVSHELL_PROJECT_ID}


clean:
	@ rm -rf bin/*

build-person-api:
	@ echo " ---         BUILDING Person API     --- "
	@ go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(PERSON_BINARY_NAME) cmd/person/main.go
	@ echo " ---      FINISH BUILD       --- "

build-job:
	@ echo " ---         BUILDING JOB       --- "
	@ go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(JOB_BINARY_NAME) cmd/jobs/main.go
	@ echo " ---      FINISH BUILD       --- "

build-crypto-api:
	@ echo " ---         BUILDING  Crypto API      --- "
	@ go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(CRYPTO_BINARY_NAME) cmd/crypto/main.go
	@ echo " ---      FINISH BUILD       --- "

set-project:
	@ sed -i "s/project/${PROJECT}/g" app.yaml

start-datastore:
	@ gcloud beta emulators datastore start --project gcp-app-engine --no-store-on-disk