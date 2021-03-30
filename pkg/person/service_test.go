package person

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-app-engine-demo/pkg/entity"
	"testing"
	"time"
)

func TestService_Store(t *testing.T) {
	repo := NewMemRepo()
	svc := NewService(repo)
	p := generatePerson()

	t.Run("store", func(t *testing.T) {
		key, err := svc.Store(p)
		assert.Nil(t, err)
		assert.Equal(t, key, p.Key)
	})

	t.Run("storeWithParentKey", func(t *testing.T) {
		p2 := &entity.Person{
			Key:       uuid.New().String(),
			FirstName: "Marcio",
			LastName:  "Cabra",
			ParentKey: p.Key,
			BirthDate: time.Now(),
		}
		key, err := svc.Store(p2)
		assert.Nil(t, err)
		assert.Equal(t, key, p2.Key)
	})

	t.Run("failToStore", func(t *testing.T) {
		p2 := &entity.Person{
			Key:       uuid.New().String(),
			FirstName: "Marcio",
			LastName:  "Cabra",
			BirthDate: time.Now(),
		}
		_, err := svc.Store(p2)
		assert.Equal(t, NewErrValidatePerson(), err)
	})
}

func TestService_FindByKeyAndFindAll(t *testing.T) {
	repo := NewMemRepo()
	svc := NewService(repo)
	p := generatePersonCollection()

	keys, _ := svc.StoreMulti(p)
	t.Run("findByKey", func(t *testing.T) {
		k, err := svc.FindByKey(keys[0])
		assert.Nil(t, err)
		assert.Equal(t, p[0].Key, k.Key)
		assert.Equal(t, p[0].FirstName, k.FirstName)

		k, err = svc.FindByKey("abc")
		assert.Equal(t, NewErrPersonNotFound(), err)
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
	svc := NewService(repo)
	p := generatePerson()
	k, _ := svc.Store(p)
	p2 := &entity.Person{
		Key:       uuid.New().String(),
		FirstName: "Marcio",
		LastName:  "Cabra",
		ParentKey: p.Key,
		BirthDate: time.Now(),
	}
	k2, _ := svc.Store(p2)

	t.Run("deleteFail", func(t *testing.T) {
		err := svc.Delete(k2)
		assert.Equal(t, NewErrDeletePerson(), err)
	})

	t.Run("delete", func(t *testing.T) {
		err := svc.Delete(k)
		assert.Nil(t, err)
		err = svc.Delete(k2)
		assert.Nil(t, err)
		_, err = svc.FindByKey(k)
		assert.Equal(t, NewErrPersonNotFound(), err)
		_, err = svc.FindByKey(k2)
		assert.Equal(t, NewErrPersonNotFound(), err)
	})
}

func generatePerson() *entity.Person {
	return &entity.Person{
		Key:       uuid.New().String(),
		FirstName: "Joaquim",
		LastName:  "Barbosa",
		BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 1, time.UTC),
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
		BirthDate: time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC),
		Address: entity.Address{
			City:  "Salvador",
			State: "BA",
		},
	}

	persons = append(persons, p, p2)
	return persons
}
