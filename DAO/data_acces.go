package DAO

import (
	"log"

	. "github.com/halil9/GoRestAPI/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CarsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "cars"
)

// Establish a connection to database
func (m *CarsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of movies
func (m *CarsDAO) FindAll() ([]Car, error) {
	var model []Cars
	err := db.C(COLLECTION).Find(bson.M{}).All(&model)
	return model, err
}

// Find a movie by its id
func (m *CarsDAO) FindById(id string) (Car, error) {
	var model Car
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&model)
	return model, err
}

// Insert a movie into database
func (m *CarsDAO) Insert(model Cars) error {
	err := db.C(COLLECTION).Insert(&model)
	return err
}

// Delete an existing movie
func (m *CarsDAO) Delete(model Car) error {
	err := db.C(COLLECTION).Remove(&model)
	return err
}

// Update an existing movie
func (m *CarsDAO) Update(model Car) error {
	err := db.C(COLLECTION).UpdateId(model.ID, &model)
	return err
}
