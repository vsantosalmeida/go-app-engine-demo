package config

import "os"

var ProjectId = os.Getenv("DATASTORE_PROJECT_ID")
var HashKey = os.Getenv("HASH_KEY")

const (
	DatastoreKind       = "Person"
	PersonApiPort       = 8080
	CryptoApiPort       = 8082
	BootstrapServers    = "bootstrap.servers"
	KafkaHost           = "localhost"
	PearsonCreatedTopic = "PERSON_CREATED_EVENT"
	SchemaId            = 1
)
