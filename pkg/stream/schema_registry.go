package stream

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-app-engine-demo/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const timeoutDuration = 3 * time.Second

type service struct {
	client *http.Client
}

func NewSchemaRegistryAPI() SchemaRegistry {
	return &service{
		client: &http.Client{},
	}
}

type subject struct {
	Subj    string `json:"subject"`
	Version int    `json:"version"`
	Id      int    `json:"id"`
}

func (s *service) GetSchemaID(subj string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.getSubjPath(subj), nil)
	if err != nil {
		log.Printf("could not create a request err:%q", err)
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("failed to get schema ID err:%q", err)
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var schema subject
		b, _ := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(b, &schema)
		if err != nil {
			log.Printf("failed do deserialize subject struct err:%q", err)
			return 0, err
		}

		return schema.Id, nil
	}

	log.Printf("request to get schema ID failed statusCode:%d", resp.StatusCode)
	return 0, errors.New("failed to get schema ID")
}

func (s *service) getSubjPath(sbj string) string {
	return fmt.Sprintf("%s/subjects/%s/versions/latest", config.SchemaRegistryHost, sbj)
}
