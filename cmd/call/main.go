package main

import (
	"context"
	"fmt"
	"log"

	"github.com/johnfercher/medium-api/internal/adapters/drivers/grpc"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.TODO()

	conn, err := grpc2.NewClient("localhost:8082", grpc2.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	getClient := grpc.NewGetProductHandlerClient(conn)
	response, err := getClient.Get(ctx, &grpc.ID{
		Id: "960b2f7e-45ec-11ef-9621-74563c34b442",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.String())
}
