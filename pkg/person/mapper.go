package person

import (
	"go-app-engine-demo/pkg/entity"
	"go-app-engine-demo/protobuf"
	"log"
	"time"
)

const dateLayout = "2006-01-02"

func mapPersonToProto(p *entity.Person) *protobuf.Person {
	address := &protobuf.Address{
		City:  p.Address.City,
		State: p.Address.State,
	}

	return &protobuf.Person{
		Key:       p.Key,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		BirthDate: p.BirthDate,
		ParentKey: p.ParentKey,
		Address:   address,
	}
}

func parsePersonBirthDate(bd string) (time.Time, error) {
	t, err := time.Parse(dateLayout, bd)
	if err != nil {
		log.Printf("Coldnt parse birth date:%s err:%q", bd, err)
		return t, err
	}

	return t, nil
}
