package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"strings"
)

// Read CSV from URL
func readCSV(csvURL string) ([][]string, error) {
	resp, err := http.Get(csvURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, err
}

// dateURLHandler expects mm-dd-yy date format and turns it into m/d/yy format
func dateURLHandler(dateFormat string) string {
	var newDate []string
	dates := strings.Split(dateFormat, "-")
	for _, date := range dates {
		newDate = append(newDate, strings.TrimLeft(date, "0"))
	}
	return strings.Join(newDate, "/")
}

// create JSON response for global/country data
func createDataJSONResponse(w http.ResponseWriter, err error, confirmed int, recovered int, deaths int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		response := &ErrorResponse{}
		response.Ok = false
		response.Error.Message = err.Error()
		json.NewEncoder(w).Encode(response)
	} else {
		response := &DataResponse{}
		response.Ok = true
		response.Data.Confirmed = confirmed
		response.Data.Recovered = recovered
		response.Data.Deaths = deaths
		response.Data.Active = response.activeCases()
		response.Data.RecoveryRate = response.recoveryRate()
		response.Data.FatalityRate = response.fatalityRate()
		json.NewEncoder(w).Encode(response)
	}
}
