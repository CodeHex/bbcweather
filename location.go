package bbcweather

import (
	_ "embed"
	"fmt"

	"github.com/codehex/bbcweather/api"
)

type Location struct {
	ID                  string
	Name                string
	Region              string
	Country             string
	Latitude, Longitude float64
}

// GetLocationByQuery returns the location for the given query. Returns the most relevant result
func GetLocationByQuery(query string) (Location, bool, error) {
	apiResults, err := api.LocationMatches(query)
	if err != nil {
		return Location{}, false, fmt.Errorf("failed to query locations API: %w", err)
	}

	if len(apiResults) == 0 {
		return Location{}, false, nil
	}

	country, ok := GetCountryFromCode(apiResults[0].Country)
	if !ok {
		return Location{}, false, fmt.Errorf("failed to get country from code: %w", err)
	}

	return Location{
		ID:        apiResults[0].ID,
		Name:      apiResults[0].Name,
		Region:    apiResults[0].Container,
		Country:   country,
		Latitude:  apiResults[0].Latitude,
		Longitude: apiResults[0].Longitude,
	}, true, nil
}
