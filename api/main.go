package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/halil9/GoRestAPI/api/DAO"
	. "github.com/halil9/GoRestAPI/api/config"
	. "github.com/halil9/GoRestAPI/api/models"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = CarsDAO{}

// create a get,post,put and delete method.
func getCars(w http.ResponseWriter, r *http.Request) {
	cars, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, cars)
}

func findCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	model, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Cars ID")
		return
	}
	respondWithJson(w, http.StatusOK, model)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var model Cars
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	model.ID = bson.NewObjectId()
	if err := dao.Insert(model); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, model)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var model Cars
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(model); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var model Cars
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(model); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
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
