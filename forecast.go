package bbcweather

import (
	"fmt"
	"time"

	"github.com/codehex/bbcweather/api"
)

type ForecastReport struct {
	ForecastLocation Location
	DayForecasts     []DayForecast
}

type DayForecast struct {
	ForecastDate        time.Time
	IssueDate           time.Time
	LastUpdated         time.Time
	WeatherDescription  string
	MaxTempC            int
	MaxTempF            int
	MinTempC            int
	MinTempF            int
	ChanceOfRainPercent int
	WindSpeedKph        int
	WindSpeedMph        int
	WindCategory        WindType
}

func GetWeatherForecast(loc Location) (ForecastReport, error) {
	apiResponse, err := api.ForecastWeather(loc.ID)
	if err != nil {
		return ForecastReport{}, fmt.Errorf("failed to get query forecast API: %w", err)
	}

	result := ForecastReport{ForecastLocation: loc}
	result.DayForecasts = []DayForecast{}
	for _, dayReport := range apiResponse.Forecasts {
		forecastDate, err := time.Parse("2006-01-02", dayReport.Summary.Report.LocalDate)
		if err != nil {
			return ForecastReport{}, fmt.Errorf("unable to convert forecast date '%s' to time: %w", dayReport.Summary.Report.LocalDate, err)
		}
		forecast := DayForecast{
			ForecastDate:        forecastDate,
			IssueDate:           dayReport.Summary.IssueDate,
			LastUpdated:         dayReport.Summary.LastUpdated,
			WeatherDescription:  dayReport.Summary.Report.WeatherTypeText,
			MaxTempC:            dayReport.Summary.Report.MaxTempC,
			MaxTempF:            dayReport.Summary.Report.MaxTempF,
			MinTempC:            dayReport.Summary.Report.MinTempC,
			MinTempF:            dayReport.Summary.Report.MinTempF,
			ChanceOfRainPercent: dayReport.Summary.Report.PrecipitationProbabilityInPercent,
			WindSpeedKph:        dayReport.Summary.Report.WindSpeedKph,
			WindSpeedMph:        dayReport.Summary.Report.WindSpeedMph,
			WindCategory:        GetWindTypeFromSpeed(dayReport.Summary.Report.WindSpeedMph),
		}
		result.DayForecasts = append(result.DayForecasts, forecast)
	}
	return result, nil
}
