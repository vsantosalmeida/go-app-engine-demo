package repository

import (
	"context"
	"go-app-engine-demo/pkg/model"
	"cloud.google.com/go/datastore"
	"log"
	"os"
)
var ctx context.Context
var projectID = os.Getenv("GCLOUD_DATASET_ID")
var datastoreClient, _ = datastore.NewClient(ctx,projectID)
const KIND = "Person"



func SavePerson(p model.Person) string {
	k := datastore.IncompleteKey(KIND, nil)
	if _, err := datastoreClient.Put(ctx, k, p); err != nil {
		log.Printf("Entity save: %s", k.String())
	}

	return k.String()
}

func QueryPersons() []*model.Person {
	q := datastore.NewQuery(KIND).Order("Firstname")

	p := make([]*model.Person, 0)
	_, _ = datastoreClient.GetAll(ctx, q, &p)

	return p
}