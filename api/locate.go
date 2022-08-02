package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// The BBC weather API using geonames to identify locations (https://www.geonames.org/). They provide
// an locator API to determine possible locations for a given query and their corresponding
// IDs that is used in the Weather API to get the weather for that location.

// Queries are made with this format
// https://open.live.bbc.co.uk/locator/locations?s=query_text&format=json&order=importance&a=true

const bbcLocationsURL = "https://open.live.bbc.co.uk/locator/locations"

type LocateAPIResponse struct {
	Response struct {
		Results struct {
			Results      []LocateAPIResult `json:"results"`
			TotalResults int               `json:"totalResults"`
		} `json:"results"`
	} `json:"response"`
}

type LocateAPIResult struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Container   string  `json:"container"`
	ContainerID int     `json:"containerId"`
	Language    string  `json:"language"`
	Timezone    string  `json:"timezone"`
	Country     string  `json:"country"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	PlaceType   string  `json:"placeType"`
}

// LocationMatches returns the matched locations that match the given query.
func LocationMatches(query string) ([]LocateAPIResult, error) {
	base, _ := url.Parse(bbcLocationsURL)
	q := url.Values{}
	q.Add("s", query)
	q.Add("format", "json")
	q.Add("order", "importance")
	q.Add("a", "true")
	base.RawQuery = q.Encode()

	resp, err := http.Get(base.String())
	if err != nil {
		return nil, fmt.Errorf("unable to query locator API: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read body from query locator API: %w", err)
	}
	var responseJson LocateAPIResponse
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return nil, fmt.Errorf("unable to decode JSON from query locator API: %w", err)
	}
	return responseJson.Response.Results.Results, nil
}
