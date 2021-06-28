package config

import (
	"log"
	"os"
)

const (
	DatastoreKind  = "Person"
	PersonApiPort  = 8080
	CryptoApiPort  = 8082
	GrpcServerHost = "localhost:9090"
	projectId      = "DATASTORE_PROJECT_ID"
	hashKey        = "HASH_KEY"
)

func GetProjectId() string {
	id := os.Getenv(projectId)
	if id == "" {
		log.Panicf("%v must not be null", projectId)
	}

	return id
}

func GetHashKey() string {
	hk := os.Getenv(hashKey)
	if hk == "" {
		log.Panicf("%v must not be null", hashKey)
	}

	return hk
}
