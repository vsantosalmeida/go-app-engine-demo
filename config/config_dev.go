package config

import "os"

var ProjectId = os.Getenv("DATASTORE_PROJECT_ID")

const (
	DatastoreKind = "Person"
	ApiPort       = 8080
)
