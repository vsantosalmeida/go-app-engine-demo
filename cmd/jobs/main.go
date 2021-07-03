package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/vsantosalmeida/go-grpc-server/protobuf"
	"go-app-engine-demo/config"
	"go-app-engine-demo/pkg/jobs"
	"go-app-engine-demo/pkg/person"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, config.GetProjectId())
	if err != nil {
		log.Fatalf("failed to create datastore client: %v", err)
	}
	defer client.Close()

	personRepo := person.NewDataStoreRepository(client)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(config.GrpcServerHost, opts...)
	if err != nil {
		log.Fatalf("failed to dial to grpc server: %v", err)
	}
	defer conn.Close()

	rpcClient := protobuf.NewPersonReceiverClient(conn)

	job := jobs.New(personRepo, rpcClient)

	start := time.Now()
	err = job.Start()
	if err != nil {
		log.Fatalf("job finish with error: %v", err)
	}
	elapsed := time.Since(start)
	log.Printf("job finished with success!\nexecution time: %s", elapsed)
}
