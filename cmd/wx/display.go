package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/codehex/bbcweather"
	"github.com/fatih/color"
)

var titleColor = color.New(color.FgWhite, color.Bold)
var labelColor = color.New(color.FgCyan, color.Bold)

func PrintLocation(loc bbcweather.Location) {
	title := loc.Name
	if loc.Region != loc.Country {
		title += " - " + loc.Region
	}
	title += " (" + loc.Country + ")"
	fmt.Println("--------------------------------------------------")
	titleColor.Println(title)
	fmt.Printf("https://www.google.com/maps/@%f,%f,14z\n", loc.Latitude, loc.Longitude)
	fmt.Println()
}

func PrintCurrentWeather(report bbcweather.CurrentWeatherReport) {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%s (updated at %s - %s)\n", titleColor.Sprint("Current Weather"), report.UpdatedAt.Format("3:04pm"), report.StationName)
	fmt.Println()
	fmt.Printf("🌡️  %s %s\n", labelColor.Sprint("Temp"), ColorizeTempC(report.TempC))
	fmt.Printf("💨 %s %s\n", labelColor.Sprint("Wind"), ColorizeWindMph(report.WindSpeedMph, report.WindCategory))
	fmt.Println()
}

func ColorizeTempC(temp int) string {
	var attr color.Attribute
	switch {
	case temp < 15:
		attr = color.FgBlue
	case temp < 25:
		attr = color.FgGreen
	case temp < 28:
		attr = color.FgYellow
	default:
		attr = color.FgRed
	}
	return color.New(attr, color.Bold).Sprintf("%d°C", temp)
}

func ColorizeWindMph(speed int, category bbcweather.WindType) string {
	var attr color.Attribute
	switch {
	case category.BeaufortNumber < 6:
		attr = color.FgGreen
	case category.BeaufortNumber < 9:
		attr = color.FgYellow
	default:
		attr = color.FgRed
	}
	return color.New(attr, color.Bold).Sprintf("%d mph (%s)", speed, category.Summary)
}

func PrintForecast(report bbcweather.ForecastReport) {
	if len(report.DayForecasts) == 0 {
		fmt.Println("No forecasts available")
		return
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("%s (updated at %s)\n", titleColor.Sprint("Forecast"), report.DayForecasts[0].LastUpdated.Format("3:04pm"))
	fmt.Println()

	w := tabwriter.NewWriter(os.Stdout, 1, 3, 3, ' ', 0)

	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
		titleColor.Sprint("Date"),
		titleColor.Sprint("High"),
		titleColor.Sprint("Low"),
		titleColor.Sprint("Wind"),
		titleColor.Sprint("Chance of Rain"),
		titleColor.Sprint("Description"))
	for _, day := range report.DayForecasts {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
			labelColor.Sprint(day.ForecastDate.Format("Mon 2 Jan")),
			ColorizeTempC(day.MaxTempC),
			ColorizeTempC(day.MinTempC),
			ColorizeWindMph(day.WindSpeedMph, day.WindCategory),
			labelColor.Sprintf("%d%%", day.ChanceOfRainPercent),
			day.WeatherDescription)

	}
	w.Flush()
}

func PrintHourlyForecast(report bbcweather.ForecastReport, day string) {
	if len(report.HourlyForecasts) == 0 {
		fmt.Println("No hourly forecasts available")
		return
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("%s (updated at %s)\n", titleColor.Sprint("Hourly Forecast"), report.DayForecasts[0].LastUpdated.Format("3:04pm"))
	fmt.Println()

	w := tabwriter.NewWriter(os.Stdout, 1, 3, 3, ' ', 0)

	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
		titleColor.Sprint("Time"),
		titleColor.Sprint("Temp"),
		titleColor.Sprint("Wind"),
		titleColor.Sprint("Description"))

	now := time.Now()
	var dayToDisplay time.Time
	switch day {
	case "today":
		dayToDisplay = now
	case "tomorrow":
		dayToDisplay = now.AddDate(0, 0, 1)
	default:
		// Should not happen
		return
	}

	for _, hour := range report.HourlyForecasts {
		if hour.ForecastDate.Day() != dayToDisplay.Day() {
			continue
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			labelColor.Sprint(hour.Timeslot),
			ColorizeTempC(hour.TemperatureC),
			ColorizeWindMph(hour.WindSpeedMph, hour.WindCategory),
			hour.Description)
	}
	w.Flush()
}
