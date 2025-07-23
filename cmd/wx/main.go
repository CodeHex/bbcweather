package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codehex/bbcweather"
)

func main() {
	today := flag.Bool("today", false, "Show today's hourly forecast")
	tomorrow := flag.Bool("tomorrow", false, "Show tomorrow's hourly forecast")
	flag.Parse()

	args := flag.Args()
	data := ""
	if len(args) > 0 {
		data = args[0]
	}

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
	if *today {
		PrintHourlyForecast(forecast, "today")
	} else if *tomorrow {
		PrintHourlyForecast(forecast, "tomorrow")
	} else {
		PrintForecast(forecast)
	}
}
