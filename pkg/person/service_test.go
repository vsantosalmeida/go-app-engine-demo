package person

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-app-engine-demo/pkg/entity"
	"testing"
)

const (
	hk                        = "xpto"
	personNotFoundInMemReason = "person not found in memory"
	personNotFoundReason      = "person not found"
)

func TestService_Store(t *testing.T) {
	repo := NewMemRepo()
	svc := NewService(repo, hk)
	p := generatePerson()

	t.Run("store", func(t *testing.T) {
		err := svc.Store(p)
		assert.Nil(t, err)
	})

	t.Run("storeWithParentKey", func(t *testing.T) {
		p2 := &entity.Person{
			Key:       uuid.New().String(),
			FirstName: "Marcio",
			LastName:  "Cabra",
			ParentKey: p.Key,
			BirthDate: "2015-07-22",
		}
		err := svc.Store(p2)
		assert.Nil(t, err)
	})

	t.Run("failToStore", func(t *testing.T) {
		p2 := &entity.Person{
			Key:       uuid.New().String(),
			FirstName: "Marcio",
			LastName:  "Cabra",
			BirthDate: "2018-05-14",
		}
		err := svc.Store(p2)
		assert.Equal(t, NewErrValidatePerson(personNotFoundReason), err)
	})

	t.Run("failToStoreUnknownParentKey", func(t *testing.T) {
		p2 := &entity.Person{
			Key:       uuid.New().String(),
			FirstName: "Marcio",
			LastName:  "Cabra",
			BirthDate: "2015-03-22",
			ParentKey: uuid.New().String(),
		}
		err := svc.Store(p2)
		assert.Equal(t, NewErrValidatePerson(personNotFoundReason), err)
	})
}

func TestService_FindByKeyAndFindAll(t *testing.T) {
	repo := NewMemRepo()
	svc := NewService(repo, hk)
	p := generatePersonCollection()
	svc.Store(p[0])
	svc.Store(p[1])

	t.Run("findByKey", func(t *testing.T) {
		k, err := svc.FindByKey(p[0].Key)
		assert.Nil(t, err)
		assert.Equal(t, p[0].Key, k.Key)
		assert.Equal(t, p[0].FirstName, k.FirstName)

		k, err = svc.FindByKey("abc")
		assert.Equal(t, NewErrPersonNotFound(personNotFoundInMemReason), err)
		assert.Nil(t, k)
	})

	t.Run("findAll", func(t *testing.T) {
		p2, err := svc.FindAll()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(p2))
	})
}

func TestService_Delete(t *testing.T) {
	repo := NewMemRepo()
	svc := NewService(repo, hk)
	p := generatePersonCollection()
	svc.Store(p[0])
	p[1].ParentKey = p[0].Key
	p[1].BirthDate = "2016-08-08"
	svc.Store(p[1])

	t.Run("deleteFail", func(t *testing.T) {
		err := svc.Delete(p[0].Key)
		assert.Equal(t, NewErrDeletePerson("person has the key associate to another person"), err)
	})

	t.Run("delete", func(t *testing.T) {
		err := svc.Delete(p[1].Key)
		assert.Nil(t, err)
		err = svc.Delete(p[0].Key)
		assert.Nil(t, err)
	})
}

func generatePerson() *entity.Person {
	return &entity.Person{
		Key:       uuid.New().String(),
		FirstName: "Joaquim",
		LastName:  "Barbosa",
		BirthDate: "1990-01-29",
		Address: entity.Address{
			City:  "São Paulo",
			State: "SP",
		},
	}
}

func generatePersonCollection() []*entity.Person {
	var persons []*entity.Person
	p := generatePerson()
	p2 := &entity.Person{
		Key:       uuid.New().String(),
		FirstName: "Maria",
		LastName:  "Souza",
		BirthDate: "1999-03-22",
		Address: entity.Address{
			City:  "Salvador",
			State: "BA",
		},
	}

	persons = append(persons, p, p2)
	return persons
}
