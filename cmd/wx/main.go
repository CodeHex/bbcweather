package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codehex/bbcweather"
)

func main() {
	args := os.Args[1:]
	data := strings.Join(args, " ")
	query := os.Getenv("WX_QUERY")
	if data != "" {
		query = data
	}
	if query == "" {
		fmt.Println("No query provided, either add a location to query as an argument or set WX_QUERY")
		os.Exit(1)
	}

	loc, found, err := bbcweather.GetLocationByQuery(query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !found {
		fmt.Println("No results found")
		os.Exit(1)
	}
	PrintLocation(loc)

	report, err := bbcweather.GetCurrentWeatherForLocation(loc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	PrintCurrentWeather(report)

	forecast, err := bbcweather.GetWeatherForecast(loc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	PrintForecast(forecast)
}
