package examples_test

import (
	"testing"

	"github.com/codehex/bbcweather"
	"github.com/matryer/is"
)

func TestLocationByTextQueryMatches(t *testing.T) {
	testData := []struct {
		Query     string
		ExpResult bbcweather.Location
	}{
		{Query: "Manchester", ExpResult: bbcweather.Location{
			ID:        "2643123",
			Name:      "Manchester",
			Region:    "Manchester",
			Country:   "United Kingdom",
			Latitude:  53.48095,
			Longitude: -2.23743,
		}},
		{Query: "Melbourne", ExpResult: bbcweather.Location{
			ID:        "2158177",
			Name:      "Melbourne",
			Region:    "Australia",
			Country:   "Australia",
			Latitude:  -37.814,
			Longitude: 144.96332,
		}},
		{Query: "Port Louis", ExpResult: bbcweather.Location{
			ID:        "934154",
			Name:      "Port Louis",
			Region:    "Mauritius",
			Country:   "Mauritius",
			Latitude:  -20.16194,
			Longitude: 57.49889,
		}},
		{Query: "Stoke Gifford", ExpResult: bbcweather.Location{
			ID:        "2636854",
			Name:      "Stoke Gifford",
			Region:    "South Gloucestershire",
			Country:   "United Kingdom",
			Latitude:  51.51686,
			Longitude: -2.54053,
		}},
	}

	for _, test := range testData {
		t.Run(test.Query, func(t *testing.T) {
			testQuery := test.Query
			testExpResult := test.ExpResult
			t.Parallel()
			is := is.New(t)
			location, found, err := bbcweather.GetLocationByQuery(testQuery)
			is.NoErr(err)                     // failed to query location API
			is.True(found)                    // location was not found
			is.Equal(location, testExpResult) // location does not match expected location
		})
	}
}

func TestLocationByTextgQueryNoMatches(t *testing.T) {
	testData := []string{"Londo", "aris", "Amsterda"}
	for _, testQuery := range testData {
		t.Run(testQuery, func(t *testing.T) {
			testQuery := testQuery
			t.Parallel()
			is := is.New(t)
			_, found, err := bbcweather.GetLocationByQuery(testQuery)
			is.NoErr(err)   // failed to query location API
			is.True(!found) // location was found
		})
	}
}
