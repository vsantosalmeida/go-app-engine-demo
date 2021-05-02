package person

import (
	"encoding/binary"
	"github.com/bearbin/go-age"
	"github.com/golang/protobuf/proto"
	"go-app-engine-demo/config"
	"go-app-engine-demo/pkg/entity"
	"go-app-engine-demo/pkg/stream"
	"go-app-engine-demo/protobuf"
)

//Service service interface
type Service struct {
	repo     Repository
	producer stream.Producer
}

//NewService create new service
func NewService(r Repository, p stream.Producer) *Service {
	return &Service{
		repo:     r,
		producer: p,
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

func (s *Service) CreateEvent(p *entity.Person) error {
	message := mapPersonToMessage(p)
	messageBytes, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	schemaIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(schemaIDBytes, uint32(config.SchemaId))

	var recordValue []byte
	recordValue = append(recordValue, byte(0))
	recordValue = append(recordValue, schemaIDBytes...)
	recordValue = append(recordValue, byte(0))
	recordValue = append(recordValue, messageBytes...)

	return s.producer.Write(recordValue, config.PearsonCreatedTopic)
}

func mapPersonToMessage(p *entity.Person) *protobuf.Person {
	address := &protobuf.Address{
		City:  p.Address.City,
		State: p.Address.State,
	}

	return &protobuf.Person{
		Key:       p.Key,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		BirthDate: p.BirthDate.Unix(),
		ParentKey: p.ParentKey,
		Address:   address,
	}
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
