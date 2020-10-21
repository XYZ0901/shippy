package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	PORT = ":50051"
)

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	Repository
}

func (s *service) CreateConsignment(ctx context.Context, in *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.Create(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.Response{
		Created:     true,
		Consignment: consignment,
	}
	return resp, nil
}

func (s *service) GetConsignments(ctx context.Context, in *pb.GetRequest) (*pb.Response, error) {
	allConsignments := s.GetAll()
	resp := &pb.Response{Consignments: allConsignments}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	log.Printf("listen on: %s\n", PORT)

	server := grpc.NewServer()
	repo := Repository{}

	pb.RegisterShippingServiceServer(server, &service{repo})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
