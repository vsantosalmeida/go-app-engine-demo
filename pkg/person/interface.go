package person

import (
	"go-app-engine-demo/pkg/entity"
)

//Reader Interface
type Reader interface {
	FindAll() ([]*entity.Person, error)
	FindByKey(k string) (*entity.Person, error)
	IsKeyAssociated(pk string) (bool, error)
}

//Writer person writer
type Writer interface {
	Store(p *entity.Person) error
	Delete(k string) error
}

//Event creation interface
type Event interface {
	CreateEvent(p *entity.Person) error
}

type Batch interface {
	// StoreMulti TODO método deve retornar algum erro em caso de falha
	StoreMulti(p []*entity.Person, success, fail chan<- *entity.Person)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
	Event
	Batch
}
