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

type dataStoreRepository struct {
	client *datastore.Client
}

func NewDataStoreRepository(c *datastore.Client) Repository {
	return &dataStoreRepository{
		client: c,
	}
}

func (r *dataStoreRepository) FindByKey(k string) (*entity.Person, error) {
	var p entity.Person
	pkey := datastore.NameKey(config.DatastoreKind, k, nil)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	log.Printf("Tying to find Person: %s", k)
	err := r.client.Get(ctx, pkey, &p)
	if err != nil {
		log.Printf("Failed to find Person: %q", err)
		return nil, NewErrPersonNotFound(err.Error())
	}

	return &p, nil
}

func (r *dataStoreRepository) isKeyAssociated(pk string) (bool, error) {
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

func (r *dataStoreRepository) FindAll() ([]*entity.Person, error) {
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

func (r *dataStoreRepository) Store(p *entity.Person) error {
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

func (r *dataStoreRepository) Delete(k string) error {
	pkey := datastore.NameKey(config.DatastoreKind, k, nil)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	log.Printf("Deleting Person: %s", k)
	err := r.client.Delete(ctx, pkey)
	if err != nil {
		log.Printf("Failed to delete Person: %q", err)
		return err
	}

	return nil
}

func (r *dataStoreRepository) Update(p *entity.Person, commitChan <-chan bool, doneChan chan<- bool) {
	var tx *datastore.Transaction
	pkey := datastore.NameKey(config.DatastoreKind, p.Key, nil)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	tx, err := r.client.NewTransaction(ctx)
	if err != nil {
		log.Printf("Failed to create a transaction: %q", err)
		doneChan <- true
		return
	}

	_, err = tx.Put(pkey, p)
	if err != nil {
		log.Printf("Failed to save Person: %q", err)
		return
	}

	commit := <-commitChan
	if commit {
		log.Printf("Commit transaction received")
		_, err = tx.Commit()
		if err != nil {
			log.Printf("Failed to commit Transaction: %q", err)
			doneChan <- true
			return
		}
		log.Printf("Entity saved: %s", p.Key)
	} else {
		log.Printf("Rollback transaction received")
		_ = tx.Rollback()
	}
	doneChan <- true
}

func (r *dataStoreRepository) GetUnsent() ([]*entity.Person, error) {
	query := datastore.NewQuery(config.DatastoreKind).Filter("Sent = ", false)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()
	var persons []*entity.Person

	_, err := r.client.GetAll(ctx, query, &persons)
	if err != nil {
		log.Printf("Failed to find unsent persons: %q", err)
		return nil, err
	}

	log.Printf("Found %d unsent persons", len(persons))
	return persons, nil
}
