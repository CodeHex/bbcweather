package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// The BBC Weather page is powered by an aggregated forecast API.

// Queries are made with this format
// https://weather-broker-cdn.api.bbci.co.uk/en/forecast/aggregated/{location_id}

const bbcForecastURL = "https://weather-broker-cdn.api.bbci.co.uk/en/forecast/aggregated/%d"

type ForecastAPIResult struct {
	Forecasts []struct {
		Detailed struct {
			IssueDate   time.Time `json:"issueDate"`
			LastUpdated time.Time `json:"lastUpdated"`
			Reports     []struct {
				EnhancedWeatherDescription        string `json:"enhancedWeatherDescription"`
				ExtendedWeatherType               int    `json:"extendedWeatherType"`
				FeelsLikeTemperatureC             int    `json:"feelsLikeTemperatureC"`
				FeelsLikeTemperatureF             int    `json:"feelsLikeTemperatureF"`
				GustSpeedKph                      int    `json:"gustSpeedKph"`
				GustSpeedMph                      int    `json:"gustSpeedMph"`
				Humidity                          int    `json:"humidity"`
				LocalDate                         string `json:"localDate"`
				PrecipitationProbabilityInPercent int    `json:"precipitationProbabilityInPercent"`
				PrecipitationProbabilityText      string `json:"precipitationProbabilityText"`
				Pressure                          int    `json:"pressure"`
				TemperatureC                      int    `json:"temperatureC"`
				TemperatureF                      int    `json:"temperatureF"`
				Timeslot                          string `json:"timeslot"`
				TimeslotLength                    int    `json:"timeslotLength"`
				Visibility                        string `json:"visibility"`
				WeatherType                       int    `json:"weatherType"`
				WeatherTypeText                   string `json:"weatherTypeText"`
				WindDescription                   string `json:"windDescription"`
				WindDirection                     string `json:"windDirection"`
				WindDirectionAbbreviation         string `json:"windDirectionAbbreviation"`
				WindDirectionFull                 string `json:"windDirectionFull"`
				WindSpeedKph                      int    `json:"windSpeedKph"`
				WindSpeedMph                      int    `json:"windSpeedMph"`
			} `json:"reports"`
		} `json:"detailed"`
		Summary struct {
			IssueDate   time.Time `json:"issueDate"`
			LastUpdated time.Time `json:"lastUpdated"`
			Report      struct {
				EnhancedWeatherDescription        string `json:"enhancedWeatherDescription"`
				GustSpeedKph                      int    `json:"gustSpeedKph"`
				GustSpeedMph                      int    `json:"gustSpeedMph"`
				LocalDate                         string `json:"localDate"`
				MaxTempC                          int    `json:"maxTempC"`
				MaxTempF                          int    `json:"maxTempF"`
				MinTempC                          int    `json:"minTempC"`
				MinTempF                          int    `json:"minTempF"`
				MostLikelyHighTemperatureC        int    `json:"mostLikelyHighTemperatureC"`
				MostLikelyHighTemperatureF        int    `json:"mostLikelyHighTemperatureF"`
				MostLikelyLowTemperatureC         int    `json:"mostLikelyLowTemperatureC"`
				MostLikelyLowTemperatureF         int    `json:"mostLikelyLowTemperatureF"`
				PollenIndex                       int    `json:"pollenIndex"`
				PollenIndexBand                   string `json:"pollenIndexBand"`
				PollenIndexIconText               string `json:"pollenIndexIconText"`
				PollenIndexText                   string `json:"pollenIndexText"`
				PollutionIndex                    int    `json:"pollutionIndex"`
				PollutionIndexBand                string `json:"pollutionIndexBand"`
				PollutionIndexIconText            string `json:"pollutionIndexIconText"`
				PollutionIndexText                string `json:"pollutionIndexText"`
				PrecipitationProbabilityInPercent int    `json:"precipitationProbabilityInPercent"`
				PrecipitationProbabilityText      string `json:"precipitationProbabilityText"`
				Sunrise                           string `json:"sunrise"`
				Sunset                            string `json:"sunset"`
				UvIndex                           int    `json:"uvIndex"`
				UvIndexBand                       string `json:"uvIndexBand"`
				UvIndexIconText                   string `json:"uvIndexIconText"`
				UvIndexText                       string `json:"uvIndexText"`
				WeatherType                       int    `json:"weatherType"`
				WeatherTypeText                   string `json:"weatherTypeText"`
				WindDescription                   string `json:"windDescription"`
				WindDirection                     string `json:"windDirection"`
				WindDirectionAbbreviation         string `json:"windDirectionAbbreviation"`
				WindDirectionFull                 string `json:"windDirectionFull"`
				WindSpeedKph                      int    `json:"windSpeedKph"`
				WindSpeedMph                      int    `json:"windSpeedMph"`
			} `json:"report"`
		} `json:"summary"`
	} `json:"forecasts"`
	IsNight     bool      `json:"isNight"`
	IssueDate   time.Time `json:"issueDate"`
	LastUpdated time.Time `json:"lastUpdated"`
	Location    struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		Container string  `json:"container"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	Night bool `json:"night"`
}

func ForecastWeather(id int) (ForecastAPIResult, error) {
	var resp *http.Response
	var err error
	for i := 0; i < 20; i++ {
		resp, err = http.Get(fmt.Sprintf(bbcForecastURL, id))
		if err != nil {
			return ForecastAPIResult{}, fmt.Errorf("unable to query forecast API: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusAccepted {
			time.Sleep(time.Second * 2)
			fmt.Println("Waiting for forecast API to process request...")
			continue
		}
		if resp.StatusCode == http.StatusOK {
			break
		}
	}

	if resp.StatusCode != http.StatusOK {
		return ForecastAPIResult{}, fmt.Errorf("unable to query forecast API after retries status code: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ForecastAPIResult{}, fmt.Errorf("unable to read body from forecast API: %w", err)
	}
	var response ForecastAPIResult
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ForecastAPIResult{}, fmt.Errorf("unable to decode JSON from forecast API: %w", err)
	}
	return response, nil
}
