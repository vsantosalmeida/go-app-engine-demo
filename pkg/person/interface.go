package person

import (
	"go-app-engine-demo/pkg/entity"
)

//reader Interface
//used to query Person in a database
type reader interface {
	FindAll() ([]*entity.Person, error)
	FindByKey(k string) (*entity.Person, error)
	isKeyAssociated(pk string) (bool, error)
}

//writer person writer
//used to save Person in a database
type writer interface {
	Store(p *entity.Person) error
	Delete(k string) error
}

//event creation interface
//used to send a message to a broker
type event interface {
	createEvent(p *entity.Person) error
}

//batch used to store a batch of Person in database
type batch interface {
	// StoreMulti TODO método deve retornar algum erro em caso de falha
	StoreMulti(p []*entity.Person, success, fail chan<- *entity.Person)
}

//encrypt interface use to log personal data of a Person with security
//must be used in crypto endpoint if want to see the content
type encrypt interface {
	encrypt(p *entity.Person) (string, error)
}

//repository repository interface
//any database system must implement these interfaces
type repository interface {
	reader
	writer
}

//UseCase use case interface
//implementation of all requirements to service layer
//control the business rules
type UseCase interface {
	reader
	writer
	event
	batch
	encrypt
}
