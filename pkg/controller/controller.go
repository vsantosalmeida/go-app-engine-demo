package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-app-engine-demo/pkg/model"
	"go-app-engine-demo/pkg/repository"
	"log"
	"net/http"
	"strings"
)

func GetPersons(w http.ResponseWriter, r *http.Request) {
	if p, err := repository.GetPersonsCollection(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print("Failed to load resources: ", err)
		return
	} else {
		log.Print("Sucess to get list")
		_ = json.NewEncoder(w).Encode(p)
	}
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	k := params["key"]
	if p, err := repository.GetPersonByKey(k); err != nil {
		http.Error(w, fmt.Sprintf("Failed to find Person with the key: %s", k), http.StatusNotFound)
		log.Printf("Failed to find Person: %s", k)
		return
	} else {
		_ = json.NewEncoder(w).Encode(p)
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	p := new(model.Person)

	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		log.Print("Failed to decode Person Entity")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if k, e := repository.SavePerson(p); e != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print("Failed to Save Person: ", err)
		return
	} else {
		formatK := strings.Split(k, ",")[1]
		_ = json.NewEncoder(w).Encode(formatK)
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	k := params["key"]
	if err := repository.DeletePerson(k); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print("Failed to delete Person: ", err)
	}

	w.WriteHeader(http.StatusNoContent)
}
