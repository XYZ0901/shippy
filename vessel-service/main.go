package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"log"
	"os"
	pb "shippy/vessel-service/proto/vessel"
)

func main() {
	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)
	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.TODO(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())
	vesselCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{vesselCollection}
	h := &handler{repository}

	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
