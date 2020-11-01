package main

import (
	"context"
	"errors"
	pb "shippy/consignment-service/proto/consignment"
	"shippy/vessel-service/proto/vessel"
)

type handler struct {
	repository
	vesselClient vessel.VesselService
}

func (s *handler) CreateConsignment(ctx context.Context,
	req *pb.Consignment, res *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vessel.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	})
	if vesselResponse == nil {
		return errors.New("error fetching vessel, returned nil")
	}
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id
	if err = s.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context,
	req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
