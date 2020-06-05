package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func main() {
	// Init gin router and Redis connection pool
	router := gin.Default()
	redisPool := newRedisPool()

	// Define routes and their handler
	api := router.Group("/api")
	{
		api.GET("/latest/global", LatestGlobalRoute(redisPool))
		api.GET("/latest/country/:countryName", LatestCountryRoute(redisPool))
		api.GET("/date/:date/global", ByDateGlobalRoute(redisPool))
		api.GET("/date/:date/country/:countryName", ByDateCountryRoute(redisPool))
	}

	// Serve
	log.Fatal(router.Run(":80"))
}

// LatestGlobalRoute responses global cases at latest date available
func LatestGlobalRoute(pool *redis.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		confirmedData, err := confirmedData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		confirmed, err := latestGlobal(confirmedData)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		recoveredData, err := recoveredData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		recovered, err := latestGlobal(recoveredData)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		deathsData, err := deathsData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		deaths, err := latestGlobal(deathsData)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		createDataJSONResponse(c, errors.Cause(err), confirmed, recovered, deaths)
	}

	return gin.HandlerFunc(fn)
}

// LatestCountryRoute responses country cases at latest date available
func LatestCountryRoute(pool *redis.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		countryName := c.Param("countryName")

		confirmedData, err := confirmedData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		confirmed, err := latestByCountry(confirmedData, countryName)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		recoveredData, err := recoveredData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		recovered, err := latestByCountry(recoveredData, countryName)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		deathsData, err := deathsData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		deaths, err := latestByCountry(deathsData, countryName)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		createDataJSONResponse(c, errors.Cause(err), confirmed, recovered, deaths)
	}

	return gin.HandlerFunc(fn)
}

// ByDateCountryRoute responses country cases at given date
func ByDateCountryRoute(pool *redis.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		countryName := c.Param("countryName")
		date := c.Param("date")

		date = dateURLHandler(date)

		confirmedData, err := confirmedData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		confirmed, err := byDateCountry(confirmedData, date, countryName)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		recoveredData, err := recoveredData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		recovered, err := byDateCountry(recoveredData, date, countryName)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		deathsData, err := deathsData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		deaths, err := byDateCountry(deathsData, date, countryName)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		createDataJSONResponse(c, errors.Cause(err), confirmed, recovered, deaths)
	}

	return gin.HandlerFunc(fn)
}

// ByDateGlobalRoute responses global cases at given date
func ByDateGlobalRoute(pool *redis.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		date := c.Param("date")

		date = dateURLHandler(date)

		confirmedData, err := confirmedData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		confirmed, err := byDateGlobal(confirmedData, date)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		recoveredData, err := recoveredData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		recovered, err := byDateGlobal(recoveredData, date)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		deathsData, err := deathsData(pool)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}
		deaths, err := byDateGlobal(deathsData, date)
		if err != nil {
			err = errors.Wrap(err, err.Error())
		}

		createDataJSONResponse(c, errors.Cause(err), confirmed, recovered, deaths)
	}

	return gin.HandlerFunc(fn)
}
