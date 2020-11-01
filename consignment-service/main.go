package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"log"
	"os"
	pb "shippy/consignment-service/proto/consignment"
	"shippy/vessel-service/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {

	Service := micro.NewService(micro.Name("go.micro.srv.consignment"))

	Service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignment")
	repository := &MongoRepository{consignmentCollection}

	vesselClient := vessel.NewVesselService("shippy.service.vessel", Service.Client())
	h := &handler{
		repository:   repository,
		vesselClient: vesselClient,
	}
	pb.RegisterShippingServiceHandler(Service.Server(), h)
	if err := Service.Run(); err != nil {
		fmt.Println(err)
	}
}
