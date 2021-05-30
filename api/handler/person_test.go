package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-app-engine-demo/pkg/entity"
	"go-app-engine-demo/pkg/person"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	hk = "xpto"
	ct = "application/json"
)

var per = &entity.Person{
	Key:       "xpto",
	FirstName: "Joaquim",
	LastName:  "Barbosa",
	BirthDate: "1990-01-29",
	Address: entity.Address{
		City:  "São Paulo",
		State: "SP",
	},
}

func TestPersonAddEndpoint(t *testing.T) {
	var tests = []struct {
		name       string
		file       string
		statusCode int
		expectErr  bool
		memRepoErr error
	}{
		{name: "When request has a correct person body must return created 201", file: "person1_200.json", statusCode: http.StatusCreated},
		{name: "When request has a person <18 with a valid parent key must return created 201", file: "person2_200.json", statusCode: http.StatusCreated},
		{name: "When request hasn't a correct person body must return bad request 400", file: "person1_400.json", statusCode: http.StatusBadRequest, expectErr: true},
		{name: "When request has a person <18 without a valid parent key must return bad request 400", file: "person2_400.json", statusCode: http.StatusBadRequest, expectErr: true},
		{name: "When request has a person with an invalid birth date key must return bad request 400", file: "person3_400.json", statusCode: http.StatusBadRequest, expectErr: true},
		{name: "When request has a person but receive an unknown err must return internal server error 500", file: "person1_200.json", statusCode: http.StatusInternalServerError, expectErr: true, memRepoErr: context.DeadlineExceeded},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := person.NewMemRepo()
			r.StubErr = tt.memRepoErr
			svc := person.NewService(r, hk)
			_ = svc.Store(per)

			ts := httptest.NewServer(personAdd(svc))
			var p *entity.Person
			defer ts.Close()

			body := bytes.NewBuffer(getFile(t, tt.file))
			res, err := http.Post(ts.URL, ct, body)
			if err != nil {
				log.Fatal(err)
			}

			b, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(b, &p)

			if tt.expectErr {
				assert.Equal(t, res.StatusCode, tt.statusCode)
				assert.Error(t, err)
				assert.Nil(t, p)
				return
			}

			assert.Equal(t, res.StatusCode, tt.statusCode)
			assert.NoError(t, err)
			assert.NotNil(t, p)
		})
	}
}

func getFile(t *testing.T, path string) []byte {
	b, err := ioutil.ReadFile(fmt.Sprintf("testData/%s", path))
	if err != nil {
		t.Errorf("Error to get file content file %q err %q", path, err)
		t.FailNow()
	}

	return b
}
