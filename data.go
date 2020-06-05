package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func confirmedData(pool *redis.Pool) ([][]string, error) {
	confirmedCSV := "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/" +
		"csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv"
	confirmed, err := readCSV(confirmedCSV, pool)
	if err != nil {
		return nil, err
	}
	return confirmed, nil
}

func recoveredData(pool *redis.Pool) ([][]string, error) {
	recoveredCSV := "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/" +
		"csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_recovered_global.csv"
	recovered, err := readCSV(recoveredCSV, pool)
	if err != nil {
		return nil, err
	}
	return recovered, nil
}

func deathsData(pool *redis.Pool) ([][]string, error) {
	deathsCSV := "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/" +
		"csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_global.csv"
	deaths, err := readCSV(deathsCSV, pool)
	if err != nil {
		return nil, err
	}
	return deaths, nil
}

// latestGlobal sums up all cases at latest date available
func latestGlobal(data [][]string) (int, error) {
	var total int
	for idx, row := range data {
		if idx > 0 {
			currentRow := row[len(row)-1]
			confirmed, err := strconv.Atoi(currentRow)
			if err != nil {
				return 0, err
			}
			total += confirmed
		}
	}
	return total, nil
}

// byDateGlobal sums up all cases at given date
func byDateGlobal(data [][]string, date string) (int, error) {
	var total int
	var dateColIndex int

	for idx, column := range data[0] {
		if date == column {
			dateColIndex = idx
		}
	}

	for idx, row := range data {
		if idx > 0 {
			currentRow := row[dateColIndex]
			confirmed, err := strconv.Atoi(currentRow)
			if err != nil {
				err := errors.New("Either date format is wrong, or data at this date does not exist")
				return 0, err
			}
			total += confirmed
		}
	}
	return total, nil
}

// latestByCountry sums up cases at latest date available within given country name
func latestByCountry(data [][]string, country string) (int, error) {
	var total int
	var countryData string
	var countryEntryFound bool = false
	regex, err := regexp.Compile("[^a-z]+")
	if err != nil {
		return 0, err
	}

	for idx, row := range data {
		countryData = regex.ReplaceAllString(strings.ToLower(row[1]), "")
		if idx > 0 && countryData == strings.ToLower(country) {
			countryEntryFound = true
			currentRow := row[len(row)-1]
			confirmed, err := strconv.Atoi(currentRow)
			if err != nil {
				return 0, err
			}
			total += confirmed
		}
	}

	if countryEntryFound == false {
		err := errors.New("Country '" + country + "' not found")
		return 0, err
	}
	return total, nil
}

// byDateCountry sums up cases at given date within given country name
func byDateCountry(data [][]string, date string, country string) (int, error) {
	var total int
	var dateColIndex int
	var countryData string
	var countryEntryFound bool = false
	regex, err := regexp.Compile("[^a-z]+")
	if err != nil {
		return 0, err
	}

	for idx, column := range data[0] {
		if date == column {
			dateColIndex = idx
		}
	}

	for idx, row := range data {
		countryData = regex.ReplaceAllString(strings.ToLower(row[1]), "")
		if idx > 0 && countryData == strings.ToLower(country) {
			countryEntryFound = true
			currentRow := row[dateColIndex]
			confirmed, err := strconv.Atoi(currentRow)
			if err != nil {
				err := errors.New("Either date format is wrong, or data at this date does not exist")
				return 0, err
			}
			total += confirmed
		}
	}

	if countryEntryFound == false {
		err := errors.New("Country '" + country + "' not found")
		return 0, err
	}

	return total, nil
}
