package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/vsantosalmeida/go-app-engine-demo/api/dto"
	"github.com/vsantosalmeida/go-app-engine-demo/pkg/entity"
	"github.com/vsantosalmeida/go-app-engine-demo/pkg/person"
)

func personAdd(service person.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p *entity.Person
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = service.Store(p)
		if err != nil {
			log.Println(err.Error())
			handleError(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err = json.NewEncoder(w).Encode(p); err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func personMultiAdd(service person.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p []*entity.Person
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var b dto.PersonBatch
		var wg sync.WaitGroup
		s := make(chan *entity.Person)
		defer close(s)
		f := make(chan *dto.FailurePerson)
		defer close(f)
		d := make(chan bool, 1)
		defer close(d)
		wg.Add(len(p))

		go service.StoreMulti(p, s, f)

		go func() {
			for {
				select {
				case v := <-s:
					b.S = append(b.S, v)
					wg.Done()
				case v := <-f:
					b.F = append(b.F, v)
					wg.Done()
				case <-d:
					return
				}
			}
		}()
		wg.Wait()
		d <- true

		w.WriteHeader(http.StatusCreated)
		if err = json.NewEncoder(w).Encode(b); err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if data == nil {
			http.Error(w, errorMessage, http.StatusNotFound)
			return
		}
		if err = json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func findPersonByKey(service person.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		k := params["key"]
		data, err := service.FindByKey(k)
		if err != nil {
			log.Println(err.Error())
			handleError(w, err)
			return
		}

		if err = json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func deletePerson(service person.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		k := params["key"]
		err := service.Delete(k)
		if err != nil {
			handleError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

func handleError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *person.ErrDeletePerson:
		http.Error(w, err.Error(), http.StatusConflict)
	case *person.ErrPersonNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
	case *person.ErrValidatePerson:
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	r.Handle("/persons", n.With(
		negroni.Wrap(personMultiAdd(service)),
	)).Methods("POST", "OPTIONS").Name("personMultiAdd")

	r.Handle("/person/{key}", n.With(
		negroni.Wrap(deletePerson(service)),
	)).Methods("DELETE", "OPTIONS").Name("deletePerson")

}
