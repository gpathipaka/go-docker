package main

import (
	"context"
	pb "go-docker/proto/consignment"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	repo IRepository
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
func (s *server) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.Response{Consignment: consignment, Created: true}, nil
}

// GetConsignments returns all the consignments
func (s *server) GetConsignments(context context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.Getall()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	log.Println("Server starting.......")
	repo := &Repository{}
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
	}
	log.Println("Server going down.......")
}
