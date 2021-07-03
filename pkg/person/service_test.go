package person

import (
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vsantosalmeida/go-app-engine-demo/pkg/entity"
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
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMemRepo()
			svc := NewService(repo, hk)
			err := svc.Store(tt.p1)
			if tt.p2 != nil {
				err = svc.Store(tt.p2)
			}
			assert.IsType(t, err, tt.expectedErr)
		})
	}
}

func TestService_FindByKey(t *testing.T) {
	var tests = []struct {
		name        string
		key         string
		expectedErr error
	}{
		{"When a key exists for a person should return it", firstPerson.Key, nil},
		{"When a key doesn't exists for a person must return err", secondPerson.Key, NewErrPersonNotFound("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMemRepo()
			svc := NewService(repo, hk)
			_ = svc.Store(firstPerson)
			p, err := svc.FindByKey(tt.key)
			assert.IsType(t, tt.expectedErr, err)
			if err == nil {
				assert.Equal(t, p, firstPerson)
			}
		})
	}
}

func TestService_Delete(t *testing.T) {
	var tests = []struct {
		name          string
		p1            *entity.Person
		p2            *entity.Person
		p1ExpectedErr error
		p2ExpectedErr error
		p1Key         string
		p2Key         string
	}{
		{name: "When delete a person which doesn't has a parent key associate to another person must delete it", p1: firstPerson, p1Key: firstPerson.Key},
		{name: "When delete a person which has a parent key associate to another person must return err", p1: firstPerson, p1Key: firstPerson.Key, p1ExpectedErr: NewErrDeletePerson(""), p2: thirdPerson, p2Key: thirdPerson.Key},
		{name: "When delete a person which the key isn't stored must return err", p1: firstPerson, p1Key: fourthPerson.Key, p1ExpectedErr: NewErrPersonNotFound("")},
		{name: "When delete a person which a parent key associate must delete it and should be possible delete the ", p1: firstPerson, p1Key: thirdPerson.Key, p2: thirdPerson, p2Key: firstPerson.Key},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMemRepo()
			svc := NewService(repo, hk)
			_ = svc.Store(tt.p1)
			if tt.p2 != nil {
				_ = svc.Store(tt.p2)
			}

			err := svc.Delete(tt.p1Key)
			assert.IsType(t, tt.p1ExpectedErr, err)

			if tt.p2 != nil {
				err = svc.Delete(tt.p2Key)
				assert.IsType(t, tt.p2ExpectedErr, err)
			}
		})
	}
}

func TestService_StoreMulti(t *testing.T) {
	var tests = []struct {
		name    string
		success int
		fail    int
		b       []*entity.Person
	}{
		{name: "When receive a batch with correct persons must store it", success: 3, fail: 0, b: []*entity.Person{firstPerson, secondPerson, thirdPerson}},
		{name: "When receive a batch with some incorrect persons must have error", success: 2, fail: 2, b: []*entity.Person{thirdPerson, secondPerson, firstPerson, fourthPerson}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(len(tt.b))
			repo := NewMemRepo()
			svc := NewService(repo, hk)
			s := make(chan *entity.Person)
			f := make(chan *entity.Person)
			d := make(chan bool)
			var totalSuccess int
			var totalFail int

			go svc.StoreMulti(tt.b, s, f)

			go func() {
				for {
					select {
					case <-s:
						totalSuccess++
						wg.Done()
					case <-f:
						totalFail++
						wg.Done()
					case <-d:
						return
					}
				}
			}()
			wg.Wait()
			d <- true

			assert.Equal(t, totalSuccess, tt.success)
			assert.Equal(t, totalFail, tt.fail)
		})
	}
}

func TestServiceIsKeyAssociated(t *testing.T) {
	var tests = []struct {
		name   string
		key    string
		expect bool
	}{
		{name: "When check an associated key of a person must return true", key: firstPerson.Key, expect: true},
		{name: "When check a not associated key of a person must return false", key: secondPerson.Key, expect: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMemRepo()
			svc := NewService(repo, hk)
			_ = svc.Store(firstPerson)
			_ = svc.Store(secondPerson)
			_ = svc.Store(thirdPerson)

			got, _ := svc.isKeyAssociated(tt.key)

			assert.Equal(t, got, tt.expect)
		})
	}
}
