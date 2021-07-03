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
	@ $(MAKE) clean
	@ go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(PERSON_BINARY_NAME) cmd/person/main.go
	@ echo " ---      FINISH BUILD       --- "

build-person-api-docker:
	@ docker build --no-cache --pull -f ./build/api/person/Dockerfile -t larolman/go-person-api .

push-person-api-docker-image:
	@ docker login
	@ docker push $(DOCKER_REPO)/go-person-api:latest

build-job:
	@ echo " ---         BUILDING JOB       --- "
	@ $(MAKE) clean
	@ go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(JOB_BINARY_NAME) cmd/jobs/main.go
	@ echo " ---      FINISH BUILD       --- "

build-job-docker:
	@ docker build --no-cache --pull -f ./build/jobs/Dockerfile -t larolman/go-job .

push-job-docker-image:
	@ docker login
	@ docker push $(DOCKER_REPO)/go-job:latest

build-crypto-api:
	@ echo " ---         BUILDING  Crypto API      --- "
	@ $(MAKE) clean
	@ go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(CRYPTO_BINARY_NAME) cmd/crypto/main.go
	@ echo " ---      FINISH BUILD       --- "

build-crypto-api-docker:
	@ docker build --no-cache --pull -f ./build/api/crypto/Dockerfile -t larolman/go-crypto-api .

push-crypto-api-docker-image:
	@ docker login
	@ docker push $(DOCKER_REPO)/go-crypto-api:latest

build-datastore-docker:
	@ docker build --no-cache -t larolman/datastore ./build/datastore-emulator

push-datastore-image:
	@ docker login
	@ docker push $(DOCKER_REPO)/datastore:latest

set-project:
	@ sed -i "s/project/${PROJECT}/g" app.yaml

start-datastore:
	@ gcloud beta emulators datastore start --project gcp-app-engine --no-store-on-disk