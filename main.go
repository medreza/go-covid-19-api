package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
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
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	confirmed, err := latestGlobal(confirmedData)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	recoveredData, err := recoveredData()
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	recovered, err := latestGlobal(recoveredData)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	deathsData, err := deathsData()
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	deaths, err := latestGlobal(deathsData)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	createDataJSONResponse(w, errors.Cause(err), confirmed, recovered, deaths)
}

// LatestCountryRoute responses country cases at latest date available
func LatestCountryRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryName := vars["countryName"]

	confirmedData, err := confirmedData()
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	confirmed, err := latestByCountry(confirmedData, countryName)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	recoveredData, err := recoveredData()
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	recovered, err := latestByCountry(recoveredData, countryName)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	deathsData, err := deathsData()
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	deaths, err := latestByCountry(deathsData, countryName)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	createDataJSONResponse(w, errors.Cause(err), confirmed, recovered, deaths)
}

// ByDateCountryRoute responses country cases at given date
func ByDateCountryRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryName := vars["countryName"]
	date := vars["date"]

	date = dateURLHandler(date)

	confirmedData, err := confirmedData()
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	confirmed, err := byDateCountry(confirmedData, date, countryName)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	recoveredData, err := recoveredData()
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	recovered, err := byDateCountry(recoveredData, date, countryName)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	deathsData, err := deathsData()
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	deaths, err := byDateCountry(deathsData, date, countryName)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	createDataJSONResponse(w, errors.Cause(err), confirmed, recovered, deaths)
}

// ByDateGlobalRoute responses global cases at given date
func ByDateGlobalRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]

	date = dateURLHandler(date)

	confirmedData, err := confirmedData()
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	confirmed, err := byDateGlobal(confirmedData, date)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	recoveredData, err := recoveredData()
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	recovered, err := byDateGlobal(recoveredData, date)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	deathsData, err := deathsData()
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}
	deaths, err := byDateGlobal(deathsData, date)
	if err != nil {
		err = errors.Wrap(err, err.Error())
	}

	createDataJSONResponse(w, errors.Cause(err), confirmed, recovered, deaths)
}
