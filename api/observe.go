package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// The BBC weather APIs provide a current weather reading based on the closest
// weather station.

// Queries are made with this format
// https://weather-broker-cdn.api.bbci.co.uk/en/observation/{location_id}

const bbcObservationURL = "https://weather-broker-cdn.api.bbci.co.uk/en/observation/%d"

type ObserveAPIResult struct {
	Station struct {
		Name     string `json:"name"`
		Distance struct {
			Km    float64 `json:"km"`
			Miles float64 `json:"miles"`
		} `json:"distance"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"station"`
	Observations []struct {
		LocalTime   string `json:"localTime"`
		LocalDate   string `json:"localDate"`
		Temperature struct {
			C int `json:"C"`
			F int `json:"F"`
		} `json:"temperature"`
		Wind struct {
			WindSpeedMph              int    `json:"windSpeedMph"`
			WindSpeedKph              int    `json:"windSpeedKph"`
			WindDirection             string `json:"windDirection"`
			WindDirectionFull         string `json:"windDirectionFull"`
			WindDirectionAbbreviation string `json:"windDirectionAbbreviation"`
		} `json:"wind"`
		HumidityPercent   int       `json:"humidityPercent"`
		PressureMb        int       `json:"pressureMb"`
		PressureDirection string    `json:"pressureDirection"`
		Visibility        string    `json:"visibility"`
		UpdateTimestamp   time.Time `json:"updateTimestamp"`
	} `json:"observations"`
}

func CurrentObservedWeather(id int) (ObserveAPIResult, error) {
	url := fmt.Sprintf(bbcObservationURL, id)

	// Try up to 3 times to allow accepted requests to be processed
	var resp *http.Response
	var err error
	for i := 0; i < 20; i++ {
		resp, err = http.Get(url)
		if err != nil {
			return ObserveAPIResult{}, fmt.Errorf("unable to query observation API: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusAccepted {
			time.Sleep(time.Second * 2)
			fmt.Println("Waiting for observation API to process request...")
			continue
		}
		if resp.StatusCode == http.StatusOK {
			break
		}
	}

	if resp.StatusCode != http.StatusOK {
		return ObserveAPIResult{}, fmt.Errorf("unable to query observation API after retires status code: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ObserveAPIResult{}, fmt.Errorf("unable to read body from observation API: %w", err)
	}

	var response ObserveAPIResult
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ObserveAPIResult{}, fmt.Errorf("unable to decode JSON from observation API: %w", err)
	}
	return response, nil
}
