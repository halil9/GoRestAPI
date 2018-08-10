package DAO

import (
	"log"

	. "github.com/halil9/GoRestAPI/api/models"
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

func (m *CarsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *CarsDAO) FindAll() ([]Cars, error) {
	var model []Cars
	err := db.C(COLLECTION).Find(bson.M{}).All(&model)
	return model, err
}

func (m *CarsDAO) FindById(id string) (Cars, error) {
	var model Car
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&model)
	return model, err
}

func (m *CarsDAO) Insert(model Cars) error {
	err := db.C(COLLECTION).Insert(&model)
	return err
}

func (m *CarsDAO) Delete(model Cars) error {
	err := db.C(COLLECTION).Remove(&model)
	return err
}

func (m *CarsDAO) Update(model Cars) error {
	err := db.C(COLLECTION).UpdateId(model.ID, &model)
	return err
}
