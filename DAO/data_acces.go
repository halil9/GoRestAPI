package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CarsDAO struct {
	Server   string
	Database string
}

var goexample *mgo.Database

const (
	COLLECTION = "cars"
)

func (m *CarsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	goexample = session.DB(m.Database)

}

func (m *CarsDAO) FindAll() ([]Car, error) {
	var cars []Car
	err := goexample.C(COLLECTION).Find(bson.M{}).All(&car)
	return cars, err
}

func (m *CarsDAO) FindById(id string) (Car, error) {
	var car Cars
	err := goexample.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&car)
	return car, err
}

func (m *CarsDAO) Insert(car Car) error {
	err := goexample.C(COLLECTION).Insert(&car)
	return err
}

func (m *CarsDAO) Delete(car Car) error {
	err := goexample.C(COLLECTION).Remove(&car)
	return err
}

func (m *CarsDAO) Update(car Car) error {
	err := goexample.C(COLLECTION).UpdateId(car.ID, &car)
	return err
}
