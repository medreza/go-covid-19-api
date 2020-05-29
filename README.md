# go-covid-19-api
[![Go Report Card](https://goreportcard.com/badge/github.com/medreza/go-covid-19-api)](https://goreportcard.com/report/github.com/medreza/go-covid-19-api)

Simple COVID-19 cases JSON API with data provided from [JHU CSSE time series *.csv files](https://github.com/CSSEGISandData/COVID-19/tree/master/csse_covid_19_data/csse_covid_19_time_series)

### API Routes
- `/api/latest/global` for latest global case summary
- `/api/latest/country/[country]` for latest specific country case summary
- `/api/date/[mm-dd-yy]/global` for global case summary at specific date
- `/api/date/[mm-dd-yy]/country/[country]` for specific country case summary at specific date

##### Notes:
- The `[mm-dd-yy]` route can be written without leading zero for single number month or date e.g. `../5-1-20/..` will return the same result as `../05-01-20/..`
- The `[country]` route is the same as country column name in the csv file, but everything is removed except alphabet. For example: route for **United Kingdom** is `../unitedkingdom` and for **Côte d'Ivoire** is `../cotedivoire`
