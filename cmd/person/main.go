package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/vsantosalmeida/go-grpc-server/protobuf"
	"go-app-engine-demo/api/handler"
	"go-app-engine-demo/config"
	"go-app-engine-demo/pkg/midleware"
	"go-app-engine-demo/pkg/person"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, config.GetProjectId())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Close()

	router := mux.NewRouter()

	personRepo := person.NewDataStoreRepository(client)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(config.GrpcServerHost, opts...)
	if err != nil {
		log.Fatalf("failed to dial to grpc server: %v", err)
	}

	rpcClient := protobuf.NewPersonReceiverClient(conn)

	personSvc := person.NewService(personRepo, rpcClient, config.GetHashKey())

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(midleware.Cors),
		negroni.NewLogger(),
	)
	//person
	handler.MakePersonHandlers(router, *n, personSvc)

	//crypto
	handler.MakeCryptoHandlers(router, *n)

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
		log.Fatal(err.Error())
	}
}
