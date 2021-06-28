package person

import (
	"context"
	"encoding/json"
	"github.com/bearbin/go-age"
	"github.com/vsantosalmeida/go-grpc-server/protobuf"
	"go-app-engine-demo/pkg/crypto"
	"go-app-engine-demo/pkg/entity"
	"google.golang.org/grpc"
	"log"
)

//service service interface
type service struct {
	repo    Repository
	client  protobuf.PersonReceiverClient
	hashKey string
}

//NewService create new service
func NewService(r Repository, client protobuf.PersonReceiverClient, hk string) UseCase {
	return &service{
		repo:    r,
		client:  client,
		hashKey: hk,
	}
}

//Store a person
func (s *service) Store(p *entity.Person) error {
	log.Print("Validating person")
	err := s.personStoreValidation(p)
	if err != nil {
		log.Print("Person validation failed")
		return err
	}

	c, err := s.encrypt(p)
	if err == nil {
		log.Printf("Saving person: %s to database", c)
	}

	err = s.repo.Store(p)
	if err == nil {
		go func() {
			log.Printf("Creating person event")
			err := s.createEvent(p)
			if err != nil {
				log.Printf("Failed to create event: %q", err)
			}
		}()
	}
	return err

}

// StoreMulti batch TODO método deve retornar algum erro em caso de falha
func (s *service) StoreMulti(p []*entity.Person, success, fail chan<- *entity.Person) {
	for _, person := range p {
		err := s.Store(person)
		if err != nil {
			fail <- person
			continue
		}
		success <- person
	}
}

// FindByKey Find a person
func (s *service) FindByKey(k string) (*entity.Person, error) {
	return s.repo.FindByKey(k)
}

//FindAll persons
func (s *service) FindAll() ([]*entity.Person, error) {
	return s.repo.FindAll()
}

func (s *service) isKeyAssociated(pk string) (bool, error) {
	return s.repo.isKeyAssociated(pk)
}

//Delete a person
func (s *service) Delete(k string) error {
	log.Printf("Deleting person key:%s", k)
	p, err := s.FindByKey(k)
	if err != nil {
		return err
	}

	a, err := s.getPersonAge(p.BirthDate)
	if err != nil {
		log.Printf("Err to get person %s age reason: %q", k, err)
		return NewErrDeletePerson(err.Error())
	}
	if a > 18 {
		// ok means the Person has a < 18 active Person
		ok, err := s.isKeyAssociated(p.Key)
		if ok || err != nil {
			log.Printf("Err to delete person, reason: person has a the key:%s associate to another person", k)
			return NewErrDeletePerson("person has the key associate to another person")
		}
	}

	return s.repo.Delete(k)
}

func (s *service) createEvent(p *entity.Person) error {
	log.Print("Event received")
	m := mapPersonToProto(p)

	var opts []grpc.CallOption
	r, err := s.client.CreateEvent(context.Background(), m, opts...)
	if err != nil {
		log.Printf("Failed to send grpc event: %q", err)
		return err
	}

	log.Printf("Success to send grpc event, reply: %v", r)
	return nil
}

func (s *service) personStoreValidation(p *entity.Person) error {
	a, err := s.getPersonAge(p.BirthDate)
	if err != nil {
		return NewErrValidatePerson(err.Error())
	}
	if a < 18 {
		log.Printf("Validating person with age less than 18")
		_, err = s.FindByKey(p.ParentKey)
		if err != nil {
			return NewErrValidatePerson("person not found")
		}
	}

	return nil
}

func (s *service) getPersonAge(bd string) (int, error) {
	t, err := parsePersonBirthDate(bd)
	if err != nil {
		return 0, err
	}

	return age.Age(t), nil
}

func (s *service) encrypt(p *entity.Person) (string, error) {
	log.Print("Encrypting person")
	data, err := json.Marshal(p)
	if err != nil {
		log.Print("Failed do marshal person")
		return "", err
	}

	c := crypto.NewCrypto(s.hashKey, data)

	err = c.Encrypt()
	if err != nil {
		log.Print("Failed to encrypt person")
		return "", err
	}
	return c.GetEncryptRaw(), nil
}
