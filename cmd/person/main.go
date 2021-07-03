package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/vsantosalmeida/go-app-engine-demo/api/handler"
	"github.com/vsantosalmeida/go-app-engine-demo/config"
	"github.com/vsantosalmeida/go-app-engine-demo/pkg/midleware"
	"github.com/vsantosalmeida/go-app-engine-demo/pkg/person"
	"github.com/vsantosalmeida/go-grpc-server/protobuf"

	"cloud.google.com/go/datastore"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, config.GetProjectId())
	if err != nil {
		log.Fatalf("failed to create datastore client: %v", err)
	}
	defer client.Close()

	router := mux.NewRouter()

	personRepo := person.NewDataStoreRepository(client)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(config.GetGrpcServerHost(), opts...)
	if err != nil {
		log.Fatalf("failed to dial to grpc server: %v", err)
	}
	defer conn.Close()

	rpcClient := protobuf.NewPersonReceiverClient(conn)

	personSvc := person.NewService(personRepo, rpcClient, config.GetHashKey())

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(midleware.Cors),
		negroni.NewLogger(),
	)
	//person
	handler.MakePersonHandlers(router, *n, personSvc)

	http.Handle("/", router)
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.PersonApiPort),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}
