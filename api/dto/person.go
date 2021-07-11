package dto

import "github.com/vsantosalmeida/go-app-engine-demo/pkg/entity"

type PersonBatch struct {
	S []*entity.Person `json:"success,omitempty"`
	F []*FailurePerson `json:"failure,omitempty"`
}

type FailurePerson struct {
	Person *entity.Person
	Reason string `json:"reason"`
}
