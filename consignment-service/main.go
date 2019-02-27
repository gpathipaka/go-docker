package main

import (
	"context"
	pb "go-docker/consignment-service/proto/consignment"
	vesselPb "go-docker/vessel-service/proto/vessel"
	"log"

	"github.com/micro/go-micro"
)

const (
	port = ":8080"
)

//IRepository is
type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	Getall() []*pb.Consignment
}

//Repository is
type Repository struct {
	consignments []*pb.Consignment
}

type server struct {
	repo         IRepository
	vesselClient vesselPb.VesselServiceClient
}

//Create is
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

//Getall returns all the items
func (repo *Repository) Getall() []*pb.Consignment {
	return repo.consignments
}

//CreateConsignment is service

func (s *server) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	//consignment, err := s.repo.Create(req)
	spec := &vesselPb.Specification{
		Capacity:  req.Weight,
		MaxWeight: int32(len(req.Containers)),
	}
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), spec)
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}

/*
	//Version 1. with only one service.
	func (s *server) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		log.Println(err)
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
} */

// GetConsignments returns all the consignments
func (s *server) GetConsignments(context context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.Getall()
	res.Consignments = consignments
	return nil
}

func main() {
	log.Println("Server starting.......")
	repo := &Repository{}
	// go-Micro Service.

	srv := micro.NewService(
		//this name must match the name given in the protobuf def.
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	//Init will parse the command line flags...
	srv.Init()

	//New Service Cleint.
	vesselClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	//Register Handler
	pb.RegisterShippingServiceHandler(srv.Server(), &server{repo, vesselClient})

	if err := srv.Run(); err != nil {
		log.Printf("Could not run the server %v", err)
	}
	// Create a new Service. Optionally include some options here.

	/*
			// GRPC Service.


		ln, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatal(err)
		}
		s := grpc.NewServer()
		pb.RegisterShippingServiceServer(s, &server{repo})

		// Register reflection service on gRPC server.
		reflection.Register(s)
		if err := s.Serve(ln); err != nil {
			log.Fatalf("Server failed to serv %v", err)
		} */
	log.Println("Server going down.......")
}
