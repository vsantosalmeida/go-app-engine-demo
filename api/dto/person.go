package dto

import "go-app-engine-demo/pkg/entity"

type PersonBatch struct {
	S []*entity.Person `json:"success,omitempty"`
	F []*entity.Person `json:"fail,omitempty"`
}
