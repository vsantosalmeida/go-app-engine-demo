package person

import "go-app-engine-demo/pkg/entity"

//Reader Interface
type Reader interface {
	FindAll() ([]*entity.Person, error)
	FindByKey(k string) (*entity.Person, error)
}

//Writer person writer
type Writer interface {
	Store(p *entity.Person) (string, error)
	StoreMulti(p []*entity.Person) ([]string, error)
	Delete(k string) error
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
}
