package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/gpathipaka/go-docker/consignment-service/proto/consignment"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

const (
	defaultFilename = "github.com/gpathipaka/go-docker/client/consignment.json"
)

func parseFile(fileName string) (*pb.Consignment, error) {
	var cons *pb.Consignment
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &cons)
	return cons, err

}

func createConsignment(client pb.ShippingServiceClient, cons *pb.Consignment) {
	res, err := client.CreateConsignment(context.TODO(), cons)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
		return
	}
	log.Println("Consignment has been created....", res.Created)
}

func getAllConsignments(client pb.ShippingServiceClient) {
	res, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Printf("Could not get the consignments.. %v", err)
	}
	log.Println(res.Consignments)
}

func main() {
	log.Println("Client started...")

	cmd.Init()
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// Contact the server and print out its response.
	//file := defaultFilename
	/* if len(os.Args) > 1 {
		file = os.Args[1]
	} */

	//cons, err := parseFile(file)

	/* if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	} */

	cons := &pb.Consignment{
		Description: "This is a test consignment",
		Weight:      550,
		Containers: []*pb.Container{
			&pb.Container{CustomerId: "cust001", UserId: "user001", Origin: "Manchester, United Kingdom"},
		},
	}
	log.Println("Calling Create Consignment with input ", cons)
	createConsignment(client, cons)

	//getAllConsignments(client)

	log.Println("Client is down...")
}
