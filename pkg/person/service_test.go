package person

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-app-engine-demo/pkg/entity"
	"testing"
)

const hk = "xpto"

var firstPerson = &entity.Person{
	Key:       uuid.New().String(),
	FirstName: "Joaquim",
	LastName:  "Barbosa",
	BirthDate: "1990-01-29",
	Address: entity.Address{
		City:  "São Paulo",
		State: "SP",
	},
}

var secondPerson = &entity.Person{
	Key:       uuid.New().String(),
	FirstName: "Maria",
	LastName:  "Souza",
	BirthDate: "1999-03-22",
	Address: entity.Address{
		City:  "Salvador",
		State: "BA",
	},
}

var thirdPerson = &entity.Person{
	Key:       uuid.New().String(),
	FirstName: "Bilbo",
	LastName:  "Bolseiro",
	ParentKey: firstPerson.Key,
	BirthDate: "2015-03-22",
	Address: entity.Address{
		City:  "Curitiba",
		State: "PR",
	},
}

var fourthPerson = &entity.Person{
	Key:       uuid.New().String(),
	FirstName: "Janaina",
	LastName:  "Errada",
	ParentKey: firstPerson.Key,
	BirthDate: "2015/03/22 12:13",
	Address: entity.Address{
		City:  "Curitiba",
		State: "PR",
	},
}

func TestService_Store(t *testing.T) {
	var tests = []struct {
		name        string
		p1          *entity.Person
		p2          *entity.Person
		expectedErr error
	}{
		{name: "When try to store person should be success", p1: firstPerson, expectedErr: nil},
		{name: "When try to store two persons wich one is age less than 18 and have a valid parent key should be success", p1: firstPerson, p2: thirdPerson, expectedErr: nil},
		{name: "When try to store person with age less than 18 and dont has a valid parentKey must return err", p1: thirdPerson, expectedErr: NewErrValidatePerson("")},
		{name: "When try to store person with an invalid birth date must return err", p1: fourthPerson, expectedErr: NewErrValidatePerson("")},
	}

	for _, tt := range tests {
		repo := NewMemRepo()
		svc := NewService(repo, hk)
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Store(tt.p1)
			if tt.p2 != nil {
				err = svc.Store(tt.p2)
			}
			assert.IsType(t, err, tt.expectedErr)
		})
	}
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

func generatePersonCollection() []*entity.Person {
	var persons []*entity.Person
	persons = append(persons, firstPerson, secondPerson)
	return persons
}
