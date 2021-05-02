package config

import "os"

var ProjectId = os.Getenv("DATASTORE_PROJECT_ID")

const (
	DatastoreKind       = "Person"
	ApiPort             = 8080
	BootstrapServers    = "bootstrap.servers"
	KafkaHost           = "localhost"
	PearsonCreatedTopic = "PERSON_CREATED_EVENT"
	SchemaId            = 1
)
