package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"go-app-engine-demo/pkg/entity"
	"go-app-engine-demo/pkg/person"
	"log"
	"net/http"
	"strings"
)

func personAdd(service person.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding Person"
		var p *entity.Person
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		var key string
		key, err = service.Store(p)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		formatKey := strings.Split(key, ",")[1]
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(formatKey); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func findAllPersons(service person.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error to find persons"
		var data []*entity.Person
		data, err := service.FindAll()
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func findPersonByKey(service person.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error to find persons"
		params := mux.Vars(r)
		k := params["key"]
		data, err := service.FindByKey(k)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func deletePerson(service person.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing person"
		params := mux.Vars(r)
		k := params["key"]
		err := service.Delete(k)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakePersonHandlers make url handlers
func MakePersonHandlers(r *mux.Router, n negroni.Negroni, service person.UseCase) {
	r.Handle("/person", n.With(
		negroni.Wrap(findAllPersons(service)),
	)).Methods("GET", "OPTIONS").Name("findAllPersons")

	r.Handle("/person/{key}", n.With(
		negroni.Wrap(findPersonByKey(service)),
	)).Methods("GET", "OPTIONS").Name("findPersonByKey")

	r.Handle("/person", n.With(
		negroni.Wrap(personAdd(service)),
	)).Methods("POST", "OPTIONS").Name("personAdd")

	r.Handle("/person/{key}", n.With(
		negroni.Wrap(deletePerson(service)),
	)).Methods("DELETE", "OPTIONS").Name("deletePerson")

}
