package handler

import (
	"bytes"
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

func TestPersonAddEndpoint(t *testing.T) {
	r := person.NewMemRepo()
	svc := person.NewService(r, hk)
	ts := httptest.NewServer(personAdd(svc))
	var p *entity.Person
	defer ts.Close()

	res, err := http.Post(ts.URL, ct, bytes.NewBuffer(getFile(t, "person_200.json")))
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &p)

	assert.NoError(t, err)
	assert.NotNil(t, p)

}

func getFile(t *testing.T, path string) []byte {
	b, err := ioutil.ReadFile(fmt.Sprintf("testData/%s", path))
	if err != nil {
		t.Errorf("Error to get file content file %q err %q", path, err)
		t.FailNow()
	}

	return b
}
