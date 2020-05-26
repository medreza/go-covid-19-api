package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// Define routes and their handler
	router.HandleFunc("/api/latest/global", LatestGlobalRoute)
	router.HandleFunc("/api/latest/{countryName}", LatestCountryRoute)
	router.HandleFunc("/api/{date}/global", ByDateGlobalRoute)
	router.HandleFunc("/api/{date}/{countryName}", ByDateCountryRoute)

	// Serve
	log.Fatal(http.ListenAndServe(":80", router))
}

// LatestGlobalRoute responses global cases at latest date available
func LatestGlobalRoute(w http.ResponseWriter, r *http.Request) {
	confirmedData, err := confirmedData()
	confirmed, err := latestGlobal(confirmedData)

	recoveredData, err := recoveredData()
	recovered, err := latestGlobal(recoveredData)

	deathsData, err := deathsData()
	deaths, err := latestGlobal(deathsData)

	createDataJSONResponse(w, err, confirmed, recovered, deaths)
}

// LatestCountryRoute responses country cases at latest date available
func LatestCountryRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryName := vars["countryName"]

	confirmedData, err := confirmedData()
	confirmed, err := latestByCountry(confirmedData, countryName)

	recoveredData, err := recoveredData()
	recovered, err := latestByCountry(recoveredData, countryName)

	deathsData, err := deathsData()
	deaths, err := latestByCountry(deathsData, countryName)

	createDataJSONResponse(w, err, confirmed, recovered, deaths)
}

// ByDateCountryRoute responses country cases at given date
func ByDateCountryRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryName := vars["countryName"]
	date := vars["date"]

	date = dateURLHandler(date)

	confirmedData, err := confirmedData()
	confirmed, err := byDateCountry(confirmedData, date, countryName)

	recoveredData, err := recoveredData()
	recovered, err := byDateCountry(recoveredData, date, countryName)

	deathsData, err := deathsData()
	deaths, err := byDateCountry(deathsData, date, countryName)

	createDataJSONResponse(w, err, confirmed, recovered, deaths)
}

// ByDateGlobalRoute responses global cases at given date
func ByDateGlobalRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]

	date = dateURLHandler(date)

	confirmedData, err := confirmedData()
	confirmed, err := byDateGlobal(confirmedData, date)

	recoveredData, err := recoveredData()
	recovered, err := byDateGlobal(recoveredData, date)

	deathsData, err := deathsData()
	deaths, err := byDateGlobal(deathsData, date)

	createDataJSONResponse(w, err, confirmed, recovered, deaths)
}
