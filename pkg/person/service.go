package person

import (
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
	return s.repo.Store(p)
}

//
func (s *Service) StoreMulti(p []*entity.Person) ([]string, error) {
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
	_, err := s.FindByKey(k)
	if err != nil {
		return err
	}
	return s.repo.Delete(k)
}
