package main

import (
	pb "github.com/gpathipaka/go-docker/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "sipping"
	vesselCollection = "vessels"
)

// Repository is
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

// VesselRepository is
type VesselRepository struct {
	session *mgo.Session
}

//FindAvailable is
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)

	if err != nil {
		return nil, err
	}

	return vessel, nil
}

//Create is
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}

//Close is
func (repo *VesselRepository) Close() {
	repo.session.Close()
}
