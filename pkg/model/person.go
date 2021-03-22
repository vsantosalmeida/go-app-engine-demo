package model

import (
	"cloud.google.com/go/datastore"
)

type Person struct {
	Key       *datastore.Key `json:"key,omitempty" datastore:"__key__"`
	Firstname string         `json:"firstname,omitempty"`
	Lastname  string         `json:"lastname,omitempty"`
	Address   Address        `json:"address,omitempty" datastore:",flatten"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
