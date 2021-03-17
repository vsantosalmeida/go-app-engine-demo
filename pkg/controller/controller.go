package controller

import (
	"encoding/json"
	"go-app-engine-demo/pkg/model"
	"go-app-engine-demo/pkg/repository"
	"net/http"
)

func GetPersons(w http.ResponseWriter, r *http.Request) {
	p := repository.QueryPersons()
	_ = json.NewEncoder(w).Encode(p)

}

func GetPerson(w http.ResponseWriter, r *http.Request) {}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p model.Person

	_ = json.NewDecoder(r.Body).Decode(&p)

	p.Key = repository.SavePerson(p)

	_ = json.NewEncoder(w).Encode(p)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {}
