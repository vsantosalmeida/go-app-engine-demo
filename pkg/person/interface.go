package person

import (
	"github.com/vsantosalmeida/go-app-engine-demo/api/dto"
	"github.com/vsantosalmeida/go-app-engine-demo/pkg/entity"
)

//Reader Interface
//used to query Person in a database
type Reader interface {
	FindAll() ([]*entity.Person, error)
	FindByKey(k string) (*entity.Person, error)
	IsKeyAssociated(pk string) (bool, error)
}

//JobReader Interface
//used from jobs to get unsent persons
type JobReader interface {
	GetUnsent() ([]*entity.Person, error)
}

//Writer person Writer
//used to save Person in a database
type Writer interface {
	Store(p *entity.Person) error
	Update(p *entity.Person, commitChan <-chan bool, doneChan chan<- bool)
	Delete(k string) error
}

//Event creation interface
//used to send a message to grpc api
type Event interface {
	CreateEvent(p *entity.Person)
}

//Batch used to store a Batch of Person in database
type Batch interface {
	StoreMulti(p []*entity.Person, success chan<- *entity.Person, failure chan<- *dto.FailurePerson)
}

//Encrypt interface use to log personal data of a Person with security
//must be used in crypto endpoint if want to see the content
type Encrypt interface {
	Encrypt(p *entity.Person) (string, error)
}

//Repository Repository interface
//any database system must implement these interfaces
type Repository interface {
	Reader
	Writer
	JobReader
}

//UseCase use case interface
//implementation of all requirements to service layer
//control the business rules
type UseCase interface {
	Reader
	Writer
	Event
	Batch
	Encrypt
}
