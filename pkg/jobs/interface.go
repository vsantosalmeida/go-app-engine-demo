package jobs

import "go-app-engine-demo/pkg/entity"

type JobService interface {
	Start() error
	getUnsent() ([]*entity.Person, error)
	send(persons []*entity.Person)
	update(person *entity.Person, commitChan <-chan bool, doneChan chan<- bool)
}
