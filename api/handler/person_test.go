package handler

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	"go-app-engine-demo/pkg/person"
	"net/http"
	"testing"
)

func TestPersonAdd(t *testing.T) {
	repo := person.NewMemRepo()
	svc := person.NewService(repo)
	r := mux.NewRouter()
	n := negroni.New()
	MakePersonHandlers(r, *n, svc)
	path, err := r.GetRoute("personAdd").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/person", path)
	payload := fmt.Sprintf(`{
		"firstname": "Chupeta",
		"lastname": "Mix",
		"address": {
			"city": "São Paulo",
			"state": "SP"
		}
}`)
	apitest.New().
		Handler(personAdd(svc)).
		Post(path).
		JSON(payload).
		Expect(t).
		Status(http.StatusCreated).
		End()
}
