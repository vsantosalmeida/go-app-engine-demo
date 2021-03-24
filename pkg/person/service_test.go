package person

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-app-engine-demo/pkg/entity"
	"testing"
)

func TestService_Store(t *testing.T) {
	repo := NewMemRepo()
	svc := NewService(repo)
	p := generatePerson()

	key, err := svc.Store(p)
	assert.Nil(t, err)
	assert.Equal(t, key, p.Key)
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
		assert.Equal(t, p[0].Firstname, k.Firstname)

		k, err = svc.FindByKey("abc")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, k)
	})

	t.Run("findAll", func(t *testing.T) {
		p, err := svc.FindAll()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(p))
	})
}

func TestService_Delete(t *testing.T) {
	repo := NewMemRepo()
	svc := NewService(repo)
	p := generatePerson()
	k, _ := svc.Store(p)

	err := svc.Delete(k)
	assert.Nil(t, err)
	_, err = svc.FindByKey(k)
	assert.Equal(t, entity.ErrNotFound, err)
}

func generatePerson() *entity.Person {
	return &entity.Person{
		Key:       uuid.New().String(),
		Firstname: "Joaquim",
		Lastname:  "Barbosa",
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
		Firstname: "Maria",
		Lastname:  "Souza",
		Address: entity.Address{
			City:  "Salvador",
			State: "BA",
		},
	}

	persons = append(persons, p, p2)
	return persons
}
