package entity

import (
	"strings"

	"cloud.google.com/go/datastore"
)

type Person struct {
	Key       string  `json:"key,omitempty" datastore:"__key__"`
	FirstName string  `json:"firstName,omitempty"`
	LastName  string  `json:"lastName,omitempty"`
	BirthDate string  `json:"birthDate,omitempty"`
	ParentKey string  `json:"parentKey,omitempty"`
	Sent      bool    `json:"sent"`
	Address   Address `json:"address,omitempty" datastore:",flatten"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// LoadKey Implement LoadKey and PropertyLoadSaver interface
func (p *Person) LoadKey(k *datastore.Key) error {
	formatKey := strings.Split(k.String(), ",")[1]
	p.Key = formatKey
	return nil
}

func (p *Person) Load(ps []datastore.Property) error {
	return datastore.LoadStruct(p, ps)
}

func (p *Person) Save() ([]datastore.Property, error) {
	return datastore.SaveStruct(p)
}
