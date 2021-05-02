package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"go-app-engine-demo/api/handler"
	"go-app-engine-demo/config"
	"go-app-engine-demo/pkg/midleware"
	"go-app-engine-demo/pkg/person"
	"go-app-engine-demo/pkg/stream"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	os.Setenv("DATASTORE_PROJECT_ID", "gcp-app-engine")
	os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:8081")

	ctx := context.Background()
	client, err := datastore.NewClient(ctx, config.ProjectId)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Close()

	router := mux.NewRouter()

	personRepo := person.NewDataStoreRepository(client, ctx)
	producer, _ := stream.NewKafkaProducer()
	personSvc := person.NewService(personRepo, producer)

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(midleware.Cors),
		negroni.NewLogger(),
	)
	//person
	handler.MakePersonHandlers(router, *n, personSvc)

	http.Handle("/", router)
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.ApiPort),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
