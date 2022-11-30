package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	Id    int
	Make  string
	Model string
	Price int
}

var vehicles = []Vehicle{
	{1, "Audi", "corolla", 100000},
	{2, "Toyota", "camry", 20000},
	{3, "Honda", "civic", 15000},
}

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func returnCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	carM := vars["make"]
	for _, car := range vehicles {
		if car.Make == carM {
			json.NewEncoder(w).Encode(car)
		}

	}
	vars = mux.Vars(r)
	carM1 := vars["model"]
	for _, car := range vehicles {
		if car.Model == carM1 {
			json.NewEncoder(w).Encode(car)
		}

	}

	vars = mux.Vars(r)
	carP, err := strconv.Atoi(vars["price"])
	if err != nil {
		fmt.Println()
	}
	for _, car := range vehicles {
		if car.Price == carP {
			json.NewEncoder(w).Encode(car)

		}

	}

	vars = mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println()
	}
	for _, car := range vehicles {
		if car.Id == carId {
			json.NewEncoder(w).Encode(car)

		}

	}

}

func updateCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("")
	}
	var updatedCar Vehicle
	json.NewDecoder(r.Body).Decode(&updatedCar)
	for _, v := range vehicles {
		if v.Id == carId {

		}

	}

	vars = mux.Vars(r)
	carP, err := strconv.Atoi(vars["price"])
	if err != nil {
		fmt.Println("")
	}
	json.NewDecoder(r.Body).Decode(&updatedCar)
	for _, v := range vehicles {
		if v.Price == carP {

		}
	}

	vars = mux.Vars(r)
	carM := vars["make"]
	json.NewDecoder(r.Body).Decode(&updatedCar)
	for _, v := range vehicles {
		if v.Make == carM {

		}
	}

	vars = mux.Vars(r)
	carM1 := vars["model"]

	json.NewDecoder(r.Body).Decode(&updatedCar)
	for _, v := range vehicles {
		if v.Model == carM1 {

		}
	}
	json.NewEncoder(w).Encode(updatedCar)

}

func createCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var newCar Vehicle
	json.NewDecoder(r.Body).Decode(&newCar)
	vehicles = append(vehicles, newCar)
	json.NewEncoder(w).Encode((vehicles))

}
func deleteCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("")
	}
	for k, v := range vehicles {
		if v.Id == carId {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
		}
	}

	vars = mux.Vars(r)
	carP, err := strconv.Atoi(vars["price"])
	if err != nil {
		fmt.Println("")
	}
	for k, v := range vehicles {
		if v.Price == carP {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
		}
	}

	vars = mux.Vars(r)
	carM := vars["make"]
	for k, v := range vehicles {
		if v.Make == carM {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
		}
	}

	vars = mux.Vars(r)
	carM1 := vars["model"]
	for k, v := range vehicles {
		if v.Model == carM1 {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
		}
	}
	json.NewEncoder(w).Encode(vehicles)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cars", returnAllCars).Methods("GET")
	router.HandleFunc("/car", createCar).Methods("POST")
	router.HandleFunc("/cars/make/{make}", returnCar).Methods("GET")
	router.HandleFunc("/cars/model/{model}", returnCar).Methods("GET")
	router.HandleFunc("/cars/price/{price}", returnCar).Methods("GET")
	router.HandleFunc("/cars/{id}", returnCar).Methods("GET")
	router.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	router.HandleFunc("/cars/model/{model}", updateCar).Methods("PUT")
	router.HandleFunc("/cars/make/{make}", updateCar).Methods("PUT")
	router.HandleFunc("/cars/price/{price}", updateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", deleteCar).Methods("DELETE")
	router.HandleFunc("/cars/price/{price}", deleteCar).Methods("DELETE")
	router.HandleFunc("/cars/make/{make}", deleteCar).Methods("DELETE")
	router.HandleFunc("/cars/model/{model}", deleteCar).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
