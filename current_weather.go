package bbcweather

import (
	"fmt"
	"time"

	"github.com/codehex/bbcweather/api"
)

type CurrentWeatherReport struct {
	ReportLocation       Location
	StationName          string
	DistFromStationKm    float64
	DistFromStationMiles float64
	UpdatedAt            time.Time

	TempC        int
	TempF        int
	WindSpeedKph int
	WindSpeedMph int
	WindCategory WindType
}

func GetCurrentWeatherForLocation(loc Location) (CurrentWeatherReport, error) {
	apiResponse, err := api.CurrentObservedWeather(loc.ID)
	if err != nil {
		return CurrentWeatherReport{}, fmt.Errorf("failed to get query observation API: %w", err)
	}

	if len(apiResponse.Observations) == 0 {
		return CurrentWeatherReport{}, fmt.Errorf("no observations found")
	}

	return CurrentWeatherReport{
		ReportLocation:       loc,
		StationName:          apiResponse.Station.Name,
		DistFromStationKm:    apiResponse.Station.Distance.Km,
		DistFromStationMiles: apiResponse.Station.Distance.Miles,
		UpdatedAt:            apiResponse.Observations[0].UpdateTimestamp,
		TempC:                apiResponse.Observations[0].Temperature.C,
		TempF:                apiResponse.Observations[0].Temperature.F,
		WindSpeedKph:         apiResponse.Observations[0].Wind.WindSpeedKph,
		WindSpeedMph:         apiResponse.Observations[0].Wind.WindSpeedMph,
		WindCategory:         GetWindTypeFromSpeed(apiResponse.Observations[0].Wind.WindSpeedMph),
	}, nil
}
