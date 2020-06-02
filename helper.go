package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Read CSV from URL with Redis caching
func readCSV(csvURL string) ([][]string, error) {
	conn := pool.Get()

	cachedCSV, _ := getRedisCache(csvURL, conn)
	if cachedCSV == "" {
		fmt.Print("[GO-COVID-19-API] CSV cache not found. Fetching from: ")
		fmt.Println(csvURL)
		resp, err := http.Get(csvURL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		cachedCSV = string(bodyBytes)
		err = setRedisCache(csvURL, cachedCSV, 12*3600, conn)
		if err != nil {
			return nil, err
		}
	}

	defer conn.Close()

	reader := csv.NewReader(strings.NewReader(cachedCSV))
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
func createDataJSONResponse(c *gin.Context, err error, confirmed int, recovered int, deaths int) {

	if err != nil {
		response := &ErrorResponse{}
		response.Ok = false
		response.Error.Message = err.Error()
		c.JSON(http.StatusOK, response)
	} else {
		response := &DataResponse{}
		response.Ok = true
		response.Data.Confirmed = confirmed
		response.Data.Recovered = recovered
		response.Data.Deaths = deaths
		response.Data.Active = response.activeCases()
		response.Data.RecoveryRate = response.recoveryRate()
		response.Data.FatalityRate = response.fatalityRate()
		c.JSON(http.StatusOK, response)
	}
}
