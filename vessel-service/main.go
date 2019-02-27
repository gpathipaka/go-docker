package main

import (
	"context"
	"errors"
	pb "go-docker/vessel-service/proto/vessel"
	"log"

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

func (v *vesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {

	return nil, errors.New("No vessel found by that spec")
}

func (s *service) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {

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
