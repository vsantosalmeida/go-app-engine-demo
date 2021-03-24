package person

import (
	"go-app-engine-demo/pkg/entity"
)

type MemRepo struct {
	m map[string]*entity.Person
}

func NewMemRepo() *MemRepo {
	var m = map[string]*entity.Person{}
	return &MemRepo{
		m: m,
	}
}

//Writer implementation
func (r *MemRepo) Store(p *entity.Person) (string, error) {
	r.m[p.Key] = p

	return p.Key, nil
}

func (r *MemRepo) StoreMulti(p []*entity.Person) ([]string, error) {
	var keys []string
	for _, person := range p {
		r.Store(person)
		keys = append(keys, person.Key)
	}

	return keys, nil
}

func (r *MemRepo) Delete(k string) error {
	if r.m[k] == nil {
		return entity.ErrNotFound
	}
	r.m[k] = nil

	return nil
}

//Reader implementation
func (r *MemRepo) FindAll() ([]*entity.Person, error) {
	var p []*entity.Person
	for _, person := range r.m {
		p = append(p, person)
	}

	return p, nil
}

func (r *MemRepo) FindByKey(k string) (*entity.Person, error) {
	if r.m[k] == nil {
		return nil, entity.ErrNotFound
	}

	return r.m[k], nil
}