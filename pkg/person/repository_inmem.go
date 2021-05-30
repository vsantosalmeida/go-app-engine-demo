package person

import (
	"go-app-engine-demo/pkg/entity"
)

type MemRepo struct {
	m       map[string]*entity.Person
	StubErr error
}

func NewMemRepo() *MemRepo {
	var m = map[string]*entity.Person{}
	return &MemRepo{
		m: m,
	}
}

// Store writer implementation
func (r *MemRepo) Store(p *entity.Person) error {
	if r.StubErr != nil {
		return r.StubErr
	}
	r.m[p.Key] = p
	return nil
}

func (r *MemRepo) Delete(k string) error {
	if r.m[k] == nil {
		return NewErrPersonNotFound("person not found in memory")
	}
	delete(r.m, k)

	return nil
}

// FindAll reader implementation
func (r *MemRepo) FindAll() ([]*entity.Person, error) {
	var p []*entity.Person
	for _, person := range r.m {
		p = append(p, person)
	}

	return p, nil
}

func (r *MemRepo) FindByKey(k string) (*entity.Person, error) {
	if r.m[k] == nil {
		return nil, NewErrPersonNotFound("person not found in memory")
	}

	return r.m[k], nil
}

func (r *MemRepo) isKeyAssociated(pk string) (bool, error) {
	for _, v := range r.m {
		if v.ParentKey == pk {
			return true, nil
		}
	}

	return false, nil
}
