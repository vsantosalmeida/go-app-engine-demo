package jobs

import (
	"context"
	"github.com/vsantosalmeida/go-grpc-server/protobuf"
	"go-app-engine-demo/pkg/entity"
	"go-app-engine-demo/pkg/person"
	"google.golang.org/grpc"
	"log"
)

type job struct {
	repo   person.Repository
	client protobuf.PersonReceiverClient
}

func New(r person.Repository, c protobuf.PersonReceiverClient) JobService {
	return &job{
		repo:   r,
		client: c,
	}
}

func (j *job) Start() error {
	log.Print("Starting new Job")
	p, err := j.getUnsent()
	if err != nil {
		log.Printf("Failed to retrive unsent persons: %q", err)
		return err
	}

	if len(p) == 0 {
		log.Print("No unsent persons find to send")
		return nil
	}

	j.send(p)

	return nil
}

func (j *job) getUnsent() ([]*entity.Person, error) {
	log.Print("Trying to find unsent persons")
	return j.repo.GetUnsent()
}

func (j *job) send(persons []*entity.Person) {
	log.Print("Sending persons")

	for _, p := range persons {
		commit := make(chan bool, 1)
		done := make(chan bool, 1)
		m := person.MapPersonToProto(p)
		p.Sent = true

		go j.repo.Update(p, commit, done)

		var opts []grpc.CallOption
		r, err := j.client.CreateEvent(context.Background(), m, opts...)
		if err != nil || !r.Created {
			log.Printf("Failed to send grpc event: %q", err)
			commit <- false
			return
		}

		commit <- true
		<-done
		log.Printf("Success to send grpc event, reply: %v", r)
	}
}

func (j *job) update(person *entity.Person, commitChan <-chan bool, doneChan chan<- bool) {
	log.Printf("Updating person: %s", person.Key)
	j.repo.Update(person, commitChan, doneChan)
}
