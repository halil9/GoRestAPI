package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func getCars(w http.ResponseWriter, r *http.Request) {
	cars, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, cars)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	car.ID = bson.NewObjectId()
	if err := dao.Insert(car); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, car)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(car); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
func deleteCar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(car); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func findCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	car, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Car ID")
		return
	}
	respondWithJson(w, http.StatusOK, car)
}

func main() {
	// router yaratıldı.
	router := mux.NewRouter()
	// router yönlendiriliyor...
	router.HandleFunc("/api/car", getCars).Methods("GET")
	router.HandleFunc("/api/car/{id}", findCar).Methods("GET")
	router.HandleFunc("/api/car/{id}", createCar).Methods("POST")
	router.HandleFunc("/api/car/{id}", updateCar).Methods("PUT")
	router.HandleFunc("/api/car/{id}", deleteCar).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}
