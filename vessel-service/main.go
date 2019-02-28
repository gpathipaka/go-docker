package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/gpathipaka/go-docker/vessel-service/proto/vessel"

	"github.com/micro/go-micro"
)

//Repository is ....
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}
type vesselRepository struct {
	vessels []*pb.Vessel
}
type service struct {
	repo Repository
}

func (repo *vesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	log.Println("FindAvailable(spec) start..")
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	log.Println("STEP 2")
	return nil, errors.New("No vessel found by that spec")
}

func (s *service) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {
	// Find the next available vessel
	log.Println("FindAvailable(ctx, in, out) start..")
	vessel, err := s.repo.FindAvailable(in)
	if err != nil {
		return err
	}
	log.Println("STEP 2", out, vessel)
	// Set the vessel as part of the response message type
	out.Vessel = vessel
	return nil
}
func main() {
	log.Println("Vessel Service started.....")
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &vesselRepository{vessels}
	s := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	s.Init()

	pb.RegisterVesselServiceHandler(s.Server(), &service{repo})

	if err := s.Run(); err != nil {
		log.Printf("Service Failed to start....")
	}
	log.Println("Vessel Service is going down.....")
}
