package main

import (
	"log"
	"os"

	pb "github.com/gpathipaka/go-docker/vessel-service/proto/vessel"

	"github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo Repository) {
	log.Println("Creating Dummy Data....")
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}

	for _, v := range vessels {
		repo.Create(v)
	}
	log.Println("exit createDummyData(repo)")
}

func main() {
	log.Println("Vessel Service started.....")

	host := os.Getenv("DB_HOST")
	log.Println("HOST", host)
	if host == "" {
		log.Println("setting DB Host to default host ", defaultHost)
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Fatalf("Error connecting to Datastore :%v", err)
	}
	defer session.Close()

	repo := &VesselRepository{session.Clone()}
	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		log.Println("The server failed to run....", err)
	}
	log.Println("Vessel Service is going down.....")
}
