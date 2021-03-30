package person

import (
	"github.com/bearbin/go-age"
	"go-app-engine-demo/pkg/entity"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Store a person
func (s *Service) Store(p *entity.Person) (string, error) {
	a := age.Age(p.BirthDate)
	if a < 18 {
		err := s.personStoreValidation(p)
		if err != nil {
			return "", err
		}
	}

	return s.repo.Store(p)
}

//Batch for store multi Persons
func (s *Service) StoreMulti(p []*entity.Person) ([]string, error) {
	for _, person := range p {
		err := s.personStoreValidation(person)
		if err != nil {
			return nil, err
		}
	}

	return s.repo.StoreMulti(p)
}

//Find a person
func (s *Service) FindByKey(k string) (*entity.Person, error) {
	return s.repo.FindByKey(k)
}

//FindAll persons
func (s *Service) FindAll() ([]*entity.Person, error) {
	return s.repo.FindAll()
}

//Delete a person
func (s *Service) Delete(k string) error {
	p, err := s.FindByKey(k)
	if err != nil {
		return err
	}

	a := age.Age(p.BirthDate)
	if a < 18 {
		_, err = s.FindByKey(p.ParentKey)
		// if error equals nil means this Person have an active parent
		if err == nil {
			return NewErrDeletePerson()
		}
	}

	return s.repo.Delete(k)
}

func (s *Service) personStoreValidation(p *entity.Person) error {
	a := age.Age(p.BirthDate)
	if a < 18 {
		_, err := s.FindByKey(p.ParentKey)
		if err != nil {
			return NewErrValidatePerson()
		}
	}

	return nil
}
