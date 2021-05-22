package person

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/google/uuid"
	"go-app-engine-demo/config"
	"go-app-engine-demo/pkg/entity"
	"google.golang.org/api/iterator"
	"log"
	"strings"
	"time"
)

const timeoutDuration = 3 * time.Second

type DataStoreRepository struct {
	client *datastore.Client
}

func NewDataStoreRepository(c *datastore.Client) *DataStoreRepository {
	return &DataStoreRepository{
		client: c,
	}
}

func (r *DataStoreRepository) FindByKey(k string) (*entity.Person, error) {
	var p entity.Person
	pkey := datastore.NameKey(config.DatastoreKind, k, nil)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	log.Printf("Tying to find Person: %s", k)
	err := r.client.Get(ctx, pkey, &p)
	if err != nil {
		log.Printf("Failed to find Person: %q", err)
		return nil, NewErrPersonNotFound()
	}

	return &p, nil
}

func (r *DataStoreRepository) IsKeyAssociated(pk string) (bool, error) {
	query := datastore.NewQuery(config.DatastoreKind).Filter("ParentKey = ", pk)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	it := r.client.Run(ctx, query)
	for {
		var person entity.Person
		_, err := it.Next(&person)
		if err == iterator.Done {
			log.Printf("No active parent key found: %q", err)
			return false, nil
		}
		if err != nil {
			log.Printf("Error fetching next person: %q", err)
			return false, err
		}
		log.Printf("Active parent key found %s", person.Key)
		return true, nil
	}
}

func (r *DataStoreRepository) FindAll() ([]*entity.Person, error) {
	var persons []*entity.Person
	q := datastore.NewQuery(config.DatastoreKind)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	log.Println("Get Person collection")
	if _, err := r.client.GetAll(ctx, q, &persons); err != nil {
		log.Printf("Failed to find Persons: %q", err)
		return nil, err
	}
	log.Printf("Found %d Persons", len(persons))
	return persons, nil
}

func (r *DataStoreRepository) Store(p *entity.Person) error {
	k := datastore.NameKey(config.DatastoreKind, uuid.New().String(), nil)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	if _, err := r.client.Put(ctx, k, p); err != nil {
		log.Printf("Failed to save Person: %q", err)
		return err
	}
	log.Printf("Entity saved: %s", k.String())
	p.Key = strings.Split(k.String(), ",")[1]

	return nil
}

func (r *DataStoreRepository) Delete(k string) error {
	client := r.client
	pkey := datastore.NameKey(config.DatastoreKind, k, nil)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	log.Printf("Deleting Person: %s", k)
	err := client.Delete(ctx, pkey)
	if err != nil {
		log.Printf("Failed to delete Person: %q", err)
		return err
	}

	return nil
}
