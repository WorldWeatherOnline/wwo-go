package main

import (
	"fmt"
	"github.com/worldweatheronline/go/wwo"
	"os"
)

func main() {
	var weather wwo.WWO

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s (apikey) (location)\n", os.Args[0])
		os.Exit(1)
	}

	weather.Key = os.Args[1]
	location := os.Args[2]

	options := map[string]string{
		"fx": "no",
	}
	forecast, err := weather.GetLocal(location, options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	fmt.Print("Current Conditions: ")

	cc := forecast.Current

	if cc.Time != 0 {
		fmt.Print("at ", cc.Time, "\n")
	}
	if cc.Temp != 0 {
		fmt.Print("Temperature\t", cc.Temp, "°C\n")
	}
	if cc.FeelsLike != 0 {
		fmt.Print("Feels Like\t", cc.FeelsLike, "°C\n")
	}
	if cc.Humidity != 0 {
		fmt.Print("Humidity\t", cc.Humidity, "%\n")
	}
	if cc.DewPoint != 0 {
		fmt.Print("Dew Point\t", cc.DewPoint, "°C\n")
	}
	if cc.Pressure != 0 {
		fmt.Print("Pressure\t", cc.Pressure, "mbar\n")
	}
	if cc.Visibility != 0 {
		fmt.Print("Visibility\t", cc.Visibility, "km\n")
	}
	if cc.CloudCover != 0 {
		fmt.Print("Cloud cover\t", cc.CloudCover, "%\n")
	}
	if cc.Precip != 0 {
		fmt.Print("Precipitation\t", cc.Precip, "mm\n")
	}
	if cc.WindSpeed != 0 {
		fmt.Print("Wind Speed\t", cc.WindSpeed, "km/h\n")
	}
	if cc.WindDir != 0 {
		fmt.Print("Wind Direction\t", cc.WindDir, "°E of N (", cc.WindDirCompass, ")\n")
	}
}
