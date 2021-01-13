package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// "github.com/go-pg/pg"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type Truck struct {
	gorm.Model

	TruckID       string      `json:"truckId"`
	DriverName    string   `json:"driverName"`
	CleanerName   string   `json:"cleanerName"`
	TruckNo       string      `json:"truckNo"`
}


var db *gorm.DB

// var err error

func main() {
	router := mux.NewRouter()

	var err error
	dataSourceName := "root:password@/golang?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Resource{})

	router.HandleFunc("/trucks", GetTrucks).Methods("GET")
	router.HandleFunc("/trucks/{id}", GetTruck).Methods("GET")
	router.HandleFunc("/trucks", CreateTruck).Methods("POST")
	router.HandleFunc("/trucks/{id}", DeleteTruck).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetTrucks(w http.ResponseWriter, r *http.Request) {
	var trucks []Truck
	db.Find(&trucks)
	json.NewEncoder(w).Encode(&trucks)
}

func GetTruck(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var truck Truck
	db.First(&truck, params["id"])
	json.NewEncoder(w).Encode(&truck)
}

func CreateTruck(w http.ResponseWriter, r *http.Request) {
	var truck Truck
	json.NewDecoder(r.Body).Decode(&truck)
	db.Create(&truck)
	json.NewEncoder(w).Encode(&truck)
}

func DeleteTruck(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var truck Truck
	db.First(&truck, params["id"])
	db.Delete(&truck)

	var trucks []Truck
	db.Find(&trucks)
	json.NewEncoder(w).Encode(&trucks)
}
