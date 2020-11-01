package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro/v2"
	"io/ioutil"
	"log"
	"os"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	ADDRESS           = "localhost:50051"
	DEFAULT_INFO_FILE = "./consignment-cli/consignment.json"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.consignment.cli"))

	service.Init()

	client := pb.NewShippingService("go.micro.srv.consignment", service.Client())

	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	resp, err := client.CreateConsignment(context.TODO(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}
	log.Printf("created: %t", resp.Created)

	resp, err = client.GetConsignments(context.TODO(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v", err)
	}
	for _, c := range resp.Consignments {
		log.Printf("%+v", c)
	}

}

func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}
