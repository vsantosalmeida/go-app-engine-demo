package person

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/google/uuid"
	"go-app-engine-demo/config"
	"go-app-engine-demo/pkg/entity"
	"log"
	"strings"
)

type DataStoreRepository struct {
	client *datastore.Client
	ctx    context.Context
}

func NewDataStoreRepository(c *datastore.Client, ctx context.Context) *DataStoreRepository {
	return &DataStoreRepository{
		client: c,
		ctx:    ctx,
	}
}

func (r *DataStoreRepository) FindByKey(k string) (*entity.Person, error) {
	var p entity.Person
	client := r.client
	pkey := datastore.NameKey(config.DatastoreKind, k, nil)
	log.Printf("Tying to find Person: %s", k)
	err := client.Get(r.ctx, pkey, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *DataStoreRepository) FindAll() ([]*entity.Person, error) {
	var persons []*entity.Person
	q := datastore.NewQuery(config.DatastoreKind)
	client := r.client
	log.Print("Get Person collection")
	if _, err := client.GetAll(r.ctx, q, &persons); err != nil {
		return nil, err
	}

	return persons, nil
}

func (r *DataStoreRepository) Store(p *entity.Person) (string, error) {
	k := datastore.NameKey(config.DatastoreKind, uuid.New().String(), nil)
	client := r.client
	if _, err := client.Put(r.ctx, k, p); err != nil {
		return "", err
	}
	log.Printf("Entity saved: %s", k.String())

	return k.String(), nil
}

func (r *DataStoreRepository) StoreMulti(p []*entity.Person) ([]string, error) {
	keys := generateKeys(len(p))
	client := r.client
	if _, err := client.PutMulti(r.ctx, keys, p); err != nil {
		return nil, err
	}
	log.Printf("Entities saved")

	return convertKeysToString(keys), nil
}

func (r *DataStoreRepository) Delete(k string) error {
	client := r.client
	pkey := datastore.NameKey(config.DatastoreKind, k, nil)

	log.Printf("Deleting Person: %s", k)
	err := client.Delete(r.ctx, pkey)
	if err != nil {
		return err
	}

	return nil
}

func generateKeys(r int) []*datastore.Key {
	var keys []*datastore.Key

	for i := 0; i < r; i++ {
		k := datastore.NameKey(config.DatastoreKind, uuid.New().String(), nil)
		keys = append(keys, k)
	}
	return keys
}

func convertKeysToString(keys []*datastore.Key) []string {
	var stringKeys []string

	for _, k := range keys {
		formatKey := strings.Split(k.String(), ",")[1]
		stringKeys = append(stringKeys, formatKey)
	}

	return stringKeys
}
