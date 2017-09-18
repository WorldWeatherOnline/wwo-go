package wwo

import "time"
import "strings"
import "strconv"
import "encoding/xml"

type Date time.Time

func (t *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}
	ti, err := time.Parse("2006-01-02", content)
	*t = Date(ti)
	return err
}

type Time12 time.Duration

func (t *Time12) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	// No moonrise, No moonset, etc.
	if strings.HasPrefix(content, "No ") {
		*t = Time12(-1)
		return nil
	}

	ti, err := time.Parse("3:04 PM", content)
	*t = Time12(ti.Sub(time.Time{}))
	return err
}

type TimeHMM time.Duration

func (t *TimeHMM) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}
	u, err := strconv.ParseUint(content, 10, 12)
	h, m := u/100, u%100
	*t = TimeHMM(time.Duration(h)*time.Hour + time.Duration(m)*time.Minute)
	return err
}

type Request struct {
	Query string `xml:"query"` // The location query used
	Type  string `xml:"type"`  // The type of location request
}

type Area struct {
	Country    string  `xml:"country"`
	Latitude   float64 `xml:"latitude"`
	Longitude  float64 `xml:"longitude"`
	Name       string  `xml:"areaName"`
	Region     string  `xml:"region"`
	Population uint    `xml:"population"`     //      Location's population
	Distance   float64 `xml:"distance_miles"` // mi   Distance between query point and this area
	WeatherURL string  `xml:"weatherUrl"`
	Zone       *Zone   `xml:"timezone"`
}

type TempRange struct {
	MaxTemp  int `xml:"maxtempC"` // °C  Maximum temperature
	MaxTempF int `xml:"maxtempF"` // °F  Maximum temperature
	MinTemp  int `xml:"mintempC"` // °C  Minimum temperature
	MinTempF int `xml:"mintempF"` // °F  Minimum temperature
}

type Weather struct {
	TempRange
	Astronomy Astronomy   `xml:"astronomy"`    // Astronomical information for the day
	Date      Date        `xml:"date"`         // Date of forecast
	SunHour   float64     `xml:"sunHour"`      // Total sun in hours
	TotalSnow float64     `xml:"totalSnow_cm"` // Total snowfall amount in cm
	UVIndex   uint        `xml:"uvIndex"`      // UV Index
	Condition []Condition `xml:"hourly"`       // Weather conditions
}

type ForecastWeather struct {
	Weather                       // includes everything in Weather
	Condition []ForecastCondition `xml:"hourly"` // Forcasted weather conditions
}

type MarineWeather struct {
	Weather                     // includes everything in Weather
	Condition []MarineCondition `xml:"hourly"`          // Forcasted weather conditions
	Tide      []Tide            `xml:"tides>tide_data"` // Tide information
}

type SkiWeather struct {
	Weather                   // includes everything in Weather
	ChanceSnow uint           `xml:"chanceofsnow"`     // %      Chance of snow
	TotalSnow  float64        `xml:"totalSnowfall_cm"` // cm   Total snowfall amount
	Top        TempRange      `xml:"top"`              // Temperature range at top
	Mid        TempRange      `xml:"mid"`              // Temperature range at middle
	Bottom     TempRange      `xml:"bottom"`           // Temperature range at bottom
	Condition  []SkiCondition `xml:"hourly"`           // Forcasted weather conditions
	Tide       []Tide         `xml:"tides>tide_data"`  // Tide information
}

type Tide struct {
	Time   Time12  `xml:"tideTime"9`     //    Local time of tide
	Height float64 `xml:"tideHeight_mt"` // m  Tide height
	Type   string  `xml:"tide_type"`     //    High, Low, Normal
}

type Astronomy struct {
	Moonrise Time12 `xml:"moonrise"` // Local time of moonrise
	Moonset  Time12 `xml:"moonset"`  // Local time of moonset
	Sunrise  Time12 `xml:"sunrise"`  // Local time of sunrise
	Sunset   Time12 `xml:"sunset"`   // Local time of sunset
}

type LevelCond struct {
	Temp              int    `xml:"tempC"`             // °C     Temperature
	TempF             int    `xml:"tempF"`             // °F     Temperature
	WindSpeed         uint   `xml:"windspeedKmph"`     // km/hr  Wind speed
	WindSpeedKnots    uint   `xml:"windspeedKnots"`    // knots  Wind speed
	WindSpeedMeterSec uint   `xml:"windspeedMeterSec"` // m/s    Wind speed
	WindSpeedMiles    uint   `xml:"windspeedMiles"`    // mi/hr  Wind speed
	WindDir           uint   `xml:"winddirDegree"`     // °EoN   Wind direction
	WindDirCompass    string `xml:"winddir16Point"`    //        Wind direction 16-point compass
	WeatherCode       uint   `xml:"weatherCode"`       //        Weather condition code <https://developer.worldweatheronline.com/api/docs/weather-icons.aspx>
	WeatherDesc       string `xml:"weatherDesc"`       //        Weather condition description
	WeatherIconUrl    string `xml:"weatherIconUrl"`    //        URL to weather icon
}

type SkiCondition struct {
	ForecastChances
	Top             LevelCond `xml:"top"`             // Temperature range at top
	Mid             LevelCond `xml:"mid"`             // Temperature range at middle
	Bottom          LevelCond `xml:"bottom"`          // Temperature range at bottom
	CloudCover      uint      `xml:"cloudcover"`      // %      Cloud cover amount
	Visibility      uint      `xml:"visibility"`      // km     Visibility
	VisibilityMiles uint      `xml:"visibilityMiles"` // mi     Visibility
	Pressure        uint      `xml:"pressure"`        // mbar   Atmospheric pressure
	PressureInches  uint      `xml:"pressureInches"`  // in     Atmospheric pressure
	Snowfall        float64   `xml:"snowfall_cm"`     // cm   Snowfall
	FreezeLevel     uint      `xml:"freezeLevel"`     // m    Freeze elevation
	Humidity        uint      `xml:"humidity"`        // %      Humidity
	Precip          float64   `xml:"precipMM"`        // mm     Precipitation
	PrecipInches    float64   `xml:"precipInches"`    // in     Precipitation
}

type Condition struct {
	Time              TimeHMM `xml:"time"`              //        Local time (Duration after start of day)
	CloudCover        uint    `xml:"cloudcover"`        // %      Cloud cover amount
	DewPoint          int     `xml:"DewPointC"`         // °C     Dew point temperature
	DewPointF         int     `xml:"DewPointF"`         // °F     Dew point temperature
	FeelsLike         int     `xml:"FeelsLikeC"`        // °C     Feels like temperature
	FeelsLikeF        int     `xml:"FeelsLikeF"`        // °F     Feels like temperature
	HeatIndex         int     `xml:"HeatIndexC"`        // °C     Heat index temperature
	HeatIndexF        int     `xml:"HeatIndexF"`        // °F     Heat index temperature
	Humidity          uint    `xml:"humidity"`          // %      Humidity
	Precip            float64 `xml:"precipMM"`          // mm     Precipitation
	PrecipInches      float64 `xml:"precipInches"`      // in     Precipitation
	Pressure          uint    `xml:"pressure"`          // mbar   Atmospheric pressure
	PressureInches    uint    `xml:"pressureInches"`    // in     Atmospheric pressure
	Temp              int     `xml:"tempC"`             // °C     Temperature
	TempF             int     `xml:"tempF"`             // °F     Temperature
	Visibility        uint    `xml:"visibility"`        // km     Visibility
	VisibilityMiles   uint    `xml:"visibilityMiles"`   // mi     Visibility
	WeatherCode       uint    `xml:"weatherCode"`       //        Weather condition code <https://developer.worldweatheronline.com/api/docs/weather-icons.aspx>
	WeatherDesc       string  `xml:"weatherDesc"`       //        Weather condition description
	WeatherIconUrl    string  `xml:"weatherIconUrl"`    //        URL to weather icon
	WindChill         int     `xml:"WindChillC"`        // °C     Wind chill temperature
	WindChillF        int     `xml:"WindChillF"`        // °F     Wind chill temperature
	WindDir           uint    `xml:"winddirDegree"`     // °EoN   Wind direction
	WindDirCompass    string  `xml:"winddir16Point"`    //        Wind direction 16-point compass
	WindGust          uint    `xml:"WindGustKmph"`      // km/hr  Wind gust
	WindGustMiles     uint    `xml:"WindGustMiles"`     // mi/hr  Wind gust
	WindSpeed         uint    `xml:"windspeedKmph"`     // km/hr  Wind speed
	WindSpeedKnots    uint    `xml:"windspeedKnots"`    // knots  Wind speed
	WindSpeedMeterSec uint    `xml:"windspeedMeterSec"` // m/s    Wind speed
	WindSpeedMiles    uint    `xml:"windspeedMiles"`    // mi/hr  Wind speed
}

type CurrentCondition struct {
	Condition
	TempF int    `xml:"temp_F"`           // °F  Temperature
	Temp  int    `xml:"temp_C"`           // °C  Temperature
	Time  Time12 `xml:"observation_time"` //     Time of the observation
}

type ForecastChances struct {
	ChanceFog      uint `xml:"chanceoffog"`      // %      Chance of fog
	ChanceFrost    uint `xml:"chanceoffrost"`    // %      Chance of front
	ChanceOvercast uint `xml:"chanceofovercast"` // %      Chance of being cloudy
	ChanceRain     uint `xml:"chanceofrain"`     // %      Chance of rain
	ChanceSnow     uint `xml:"chanceofsnow"`     // %      Chance of snow
	ChanceHighTemp uint `xml:"chanceofhightemp"` // %      Chance of high temperatures FIXME not in docs
	ChanceDry      uint `xml:"chanceofremdry"`   // %      Chance of remaining dry FIXME not in docs
	ChanceSunshine uint `xml:"chanceofsunshine"` // %      Chance of being sunny
	ChanceThunder  uint `xml:"chanceofthunder"`  // %      Chance of thunder and/or lightning
	ChanceWindy    uint `xml:"chanceofwindy"`    // %      Chance of being windy
}

type ForecastCondition struct {
	Condition
	ForecastChances
}

type MarineCondition struct {
	Condition
	SigHeight       float64 `xml:"sigHeight_m"`      // m    Significant wave height
	SwellHeight     float64 `xml:"swellHeight_m"`    // m    Swell wave height
	SwellHeight_ft  float64 `xml:"swellHeight_ft"`   // ft   Swell wave height FIXME docs say swell_Height_ft
	SwellDir        uint    `xml:"swellDir"`         // °EoN Swell direction
	SwellDirCompass string  `xml:"swellDir16Point"`  //      Swell compass direction
	SwellPeriod     float64 `xml:"swellPeriod_secs"` // sec  Swell period
	WaterTemp       int     `xml:"waterTemp_C"`      // °C   Water temperature
	WaterTemp_F     int     `xml:"waterTemp_F"`      // °F   Water temperature
}

type ClimateAverage struct {
	Index                uint    `xml:"index"`                   //        Month index Integer: 1-12
	Name                 string  `xml:"name"`                    //        The name of the month
	MinTemp              float64 `xml:"avgMinTemp"`              // °C     Average minimum temperature
	MinTemp_F            float64 `xml:"avgMinTemp_F"`            // °F     Average minimum temperature
	MaxTemp              float64 `xml:"avgMaxTemp"`              // °C     Average maximum temperature
	MaxTemp_F            float64 `xml:"avgMaxTemp_F"`            // °F     Average maximum temperature
	AbsMinTemp           float64 `xml:"absMinTemp"`              // °C     Absolute minimum temperature
	AbsMinTemp_F         float64 `xml:"absMinTemp_F"`            // °F     Absolute minimum temperature
	AbsMaxTemp           float64 `xml:"absMaxTemp"`              // °C     Absolute maximum temperature
	AbsMaxTemp_F         float64 `xml:"absMaxTemp_F"`            // °F     Absolute maximum temperature
	Temp                 float64 `xml:"avgTemp"`                 // °C     Average temperature
	Temp_F               float64 `xml:"avgTemp_F"`               // °F     Average temperature
	MaxWindSpeed         float64 `xml:"maxWindSpeed_kmph"`       // km/hr  Maximum wind speed FIXME average or absolute?
	MaxWindSpeed_mph     float64 `xml:"maxWindSpeed_mph"`        // mi/hr  Maximum wind speed
	MaxWindSpeed_knots   float64 `xml:"maxWindSpeed_knots"`      // knots  Maximum wind speed
	MaxWindSpeed_ms      float64 `xml:"maxWindSpeed_ms"`         // m/s    Maximum wind speed
	WindSpeed            float64 `xml:"avgWindSpeed_kmph"`       // km/hr  Average wind speed
	WindSpeed_miles      float64 `xml:"avgWindSpeed_miles"`      // mi/hr  Average wind speed
	WindSpeed_knots      float64 `xml:"avgWindSpeed_knots"`      // knots  Average wind speed
	WindSpeed_ms         float64 `xml:"avgWindSpeed_ms"`         // m/s    Average wind speed
	WindGust             float64 `xml:"avgWindGust_kmph"`        // km/hr  Average wind gust
	WindGust_miles       float64 `xml:"avgWindGust_miles"`       // mi/hr  Average wind gust
	WindGust_knots       float64 `xml:"avgWindGust_knots"`       // knots  Average wind gust
	WindGust_ms          float64 `xml:"avgWindGust_ms"`          // m/s    Average wind gust
	DailyRainfall        float64 `xml:"avgDailyRainfall"`        // mm     Average daily rainfall
	DailyRainfall_inch   float64 `xml:"avgDailyRainfall_inch"`   // in     Average daily rainfall
	MonthlyRainfall      float64 `xml:"avgMonthlyRainfall"`      // mm     Average monthly rainfall
	MonthlyRainfall_inch float64 `xml:"avgMonthlyRainfall_inch"` // in     Average monthly rainfall
	Humidity             float64 `xml:"avgHumidity"`             // %      Average humidity
	Cloud                float64 `xml:"avgCloud"`                // %      Average cloud cover
	Visibility           float64 `xml:"avgVis_km"`               // km     Average visibility
	Visibility_miles     float64 `xml:"avgVis_miles"`            // mi     Average visibility
	Pressure             float64 `xml:"avgPressure_mb"`          // mbar   Average pressure
	Pressure_inch        float64 `xml:"avgPressure_inch"`        // in     Average pressure
	DryDays              uint    `xml:"avgDryDays"`              //        Average number of dry days
	RainDays             uint    `xml:"avgRainDays"`             //        Average number of rain days
	SnowDays             uint    `xml:"avgSnowDays"`             //        Average number of snow days
	FogDays              uint    `xml:"avgFogDays"`              //        Average number of foggy days
	ThunderDays          uint    `xml:"avgThunderDays"`          //        Average number of thunder days
	UVIndex              uint    `xml:"avgUVIndex"`              //        Average UV Index
	SunHour              float64 `xml:"avgSunHour"`              // hr/day Average Sun
}

type Zone struct {
	Offset float64 `xml:"utcOffset"` // hr  Offset from UTC including fractional hours
}

type Local struct {
	Area    Area              `xml:"nearest_area"`
	Climate []ClimateAverage  `xml:"ClimateAverages>month"`
	Current CurrentCondition  `xml:"current_condition"`
	Request Request           `xml:"request"`
	Weather []ForecastWeather `xml:"weather"`
}

type Marine struct {
	Request Request         `xml:"request"`
	Area    Area            `xml:"nearest_area"`
	Weather []MarineWeather `xml:"weather"`
}

type PastLocal struct {
	Request Request   `xml:"request"`
	Area    Area      `xml:"nearest_area"`
	Weather []Weather `xml:"weather"`
}

type PastMarine Marine

type Ski struct {
	Request Request      `xml:"request"`
	Area    Area         `xml:"nearest_area"`
	Weather []SkiWeather `xml:"weather"`
}

type TimeZone struct {
	Request Request `xml:"request"`
	Area    Area    `xml:"nearest_area"`
	Zone    Zone    `xml:"time_zone"`
}

type Search struct {
	Area []Area `xml:"result"`
}
