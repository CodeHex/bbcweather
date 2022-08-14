package examples_test

import (
	"testing"

	"github.com/codehex/bbcweather"
	"github.com/matryer/is"
)

func TestCurrentWeatherReport(t *testing.T) {
	is := is.New(t)
	testLoc := bbcweather.Location{
		ID:        "2643123",
		Name:      "Manchester",
		Region:    "Manchester",
		Country:   "United Kingdom",
		Latitude:  53.48095,
		Longitude: -2.23743,
	}

	report, err := bbcweather.GetCurrentWeatherForLocation(testLoc)
	is.NoErr(err) // failed to get current weather report from API
	is.Equal("2643123", report.ReportLocation.ID)
	is.Equal("Rostherne No 2", report.StationName)
	is.Equal(9.24, report.DistFromStationMiles)
	is.Equal(14.87, report.DistFromStationKm)
}
