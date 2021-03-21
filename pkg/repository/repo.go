package repository

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/google/uuid"
	"go-app-engine-demo/pkg/model"
	"log"
	"os"
)

var projectID = os.Getenv("DATASTORE_PROJECT_ID")

const KIND = "Person"

func SavePerson(p *model.Person) (string, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	k := datastore.NameKey(KIND, uuid.New().String(), nil)
	if _, err := client.Put(ctx, k, p); err != nil {
		return "", err
	}
	log.Printf("Entity saved: %s", k.String())

	return k.String(), nil
}

func GetPersonByKey(k string) (model.Person, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var p model.Person
	pkey := datastore.NameKey(KIND, k, nil)

	log.Printf("Tying to find Person: %s", k)
	err = client.Get(ctx, pkey, &p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func DeletePerson(k string) error {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	pkey := datastore.NameKey(KIND, k, nil)

	log.Printf("Deleting Person: %s", k)
	err = client.Delete(ctx, pkey)
	if err != nil {
		return err
	}

	return nil
}

func GetPersonsCollection() ([]model.Person, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	q := datastore.NewQuery(KIND)
	var persons []model.Person

	log.Print("Get Person collection")
	if _, err = client.GetAll(ctx, q, &persons); err != nil {
		return nil, err
	}

	return persons, nil
}
