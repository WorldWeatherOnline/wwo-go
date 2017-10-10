# worldweatheronline

This is a Go library for the worldweatheronline.com APIs.

## Installing

```sh
go get github.com/worldweatheronline/wwo-go/wwo
```

## Usage

For more details see the example application in [example/main.go](https://github.com/WorldWeatherOnline/wwo-go/blob/master/example/main.go) and [documentation](https://godoc.org/github.com/WorldWeatherOnline/wwo-go/wwo).

```go
var weather = WWO({"your-hex-api-key-goes-in-here!"})
forecast, err := weather.GetLocal("London", map[string]string{})
if err == nil {
	fmt.Print("Current Temperature: ", forecast.Current.Temp, "°C\n")
}
```

### Example Output

```
$ ./example (key) London
Current Conditions: at 14:45
Temperature	20°C
Feels Like	20°C
Humidity	60%
Pressure	1022mbar
Visibility	10km
Wind Speed	15km/h
Wind Direction	90°E of N (E)
```

