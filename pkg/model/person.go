package model

import (
	"cloud.google.com/go/datastore"
	"strings"
)

type Person struct {
	Key       string  `json:"key,omitempty"`
	Firstname string  `json:"firstname,omitempty"`
	Lastname  string  `json:"lastname,omitempty"`
	Address   Address `json:"address,omitempty" datastore:",flatten"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

func (p *Person) LoadKey(k *datastore.Key) error {
	splitKey := strings.Split(k.String(), ",")
	p.Key = splitKey[1]
	return nil
}

func (p *Person) Load(ps []datastore.Property) error {
	return datastore.LoadStruct(p, ps)
}

func (p *Person) Save() ([]datastore.Property, error) {
	return datastore.SaveStruct(p)
}
