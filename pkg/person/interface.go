package person

import (
	"go-app-engine-demo/pkg/entity"
)

//reader Interface
type reader interface {
	FindAll() ([]*entity.Person, error)
	FindByKey(k string) (*entity.Person, error)
	IsKeyAssociated(pk string) (bool, error)
}

//writer person writer
type writer interface {
	Store(p *entity.Person) error
	Delete(k string) error
}

//event creation interface
type event interface {
	CreateEvent(p *entity.Person) error
}

type batch interface {
	// StoreMulti TODO método deve retornar algum erro em caso de falha
	StoreMulti(p []*entity.Person, success, fail chan<- *entity.Person)
}

type encrypt interface {
	encrypt(p *entity.Person) (string, error)
}

//repository repository interface
type repository interface {
	reader
	writer
}

//UseCase use case interface
type UseCase interface {
	reader
	writer
	event
	batch
	encrypt
}
