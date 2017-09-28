/*
wwo provides an interface to the premium api of worldweatheronline.com

This requires an API key, held by a WWO structure,
which is then used to perform queries.

 var weather = WWO({"your-hex-api-key-goes-in-here!"})
 forecast, err := weather.GetLocal("London", map[string]string{})

The optional options passed in the map are documented with the various Get functions.
Each Get function returns a structure of the appropriate type and a possible error.
That error will be set for any transport, unmashalling, or API errors,
depending on the type of error, including all API errors, the structure may also be filled in to some extent.

*/
package wwo

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Essential information for WorldWeatherOnline lookups.
type WWO struct {
	Key      string // API key
	Insecure bool   // Use http rather than https
}

func (w *WWO) fetch(service string, query map[string]string) ([]byte, error) {
	var u url.URL

	if w.Insecure {
		u.Scheme = "http"
	} else {
		u.Scheme = "https"
	}

	u.Host = "api.worldweatheronline.com"
	u.Path = "/premium/v1/" + service + ".ashx"

	var values = make(url.Values)

	values.Set("key", w.Key)
	for k, v := range query {
		values.Set(k, v)
	}
	values.Set("format", "xml")
	u.RawQuery = values.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return text, nil
}

// Fetch a local forecast for location.
//
// Supported options are (defaults marked with *):
//   num_of_days      Number of days of forecast to include (0-21, *14)
//   date             Start date of forecast (today, *tomorrow, YYYY-mm-dd)
//   fx               Include forecast (*yes, no)
//   cc               Include current conditions (*yes, no)
//   mca              Include monthly averages (*yes, no)
//   fx24             Include tp-hourly forecasts (*yes, no)
//   includelocation  Include nearest location information (yes, *no)
//   tp               Number of hours in detailed forecast (1, *3, 6, 12, 24)
func (w *WWO) GetLocal(location string, opt map[string]string) (*Local, error) {
	opt["q"] = location
	opt["date_format"] = ""

	text, err := w.fetch("weather", opt)
	if err != nil {
		return nil, err
	}

	var o *Local = new(Local)
	err = xml.Unmarshal(text, o)
	if err != nil {
		return o, err
	}

	if o.Error != nil {
		return o, errors.New(*o.Error)
	}

	return o, nil
}

// Fetch a marine forecast for location.
//
// Supported options are (defaults marked with *):
//   fx    Include forecast (*yes, no)
//   tp    Number of hours in detailed forecast (1, *3, 6, 12, 24)
//   tide  Include tide information (yes, *no)
func (w *WWO) GetMarine(location string, opt map[string]string) (*Marine, error) {
	opt["q"] = location
	opt["date_format"] = ""

	text, err := w.fetch("marine", opt)
	if err != nil {
		return nil, err
	}

	var o *Marine = new(Marine)
	err = xml.Unmarshal(text, o)
	if err != nil {
		return o, err
	}

	if o.Error != nil {
		return o, errors.New(*o.Error)
	}

	return o, nil
}

// Fetch a ski forecast for location.
//
// Supported options are (defaults marked with *):
//   num_of_days      Number of days of forecast to include (0-21, *14)
//   date             Start date of forecast (today, *tomorrow, YYYY-mm-dd)
//   includelocation  Include nearest location information (yes, *no)
func (w *WWO) GetSki(location string, opt map[string]string) (*Ski, error) {
	opt["q"] = location
	opt["date_format"] = ""

	text, err := w.fetch("ski", opt)
	if err != nil {
		return nil, err
	}

	var o *Ski = new(Ski)
	err = xml.Unmarshal(text, o)
	if err != nil {
		return o, err
	}

	if o.Error != nil {
		return o, errors.New(*o.Error)
	}

	return o, nil
}

// Fetch historical local weather information for location.
//
// Supported options are (defaults marked with *):
//   date             Start date (YYYY-mm-dd)
//   enddate          End date (YYYY-mm-dd)
//   includelocation  Include nearest location information (yes, *no)
//   tp               Number of hours in detailed forecast (1, *3, 6, 12, 24)
func (w *WWO) GetPastLocal(location string, opt map[string]string) (*PastLocal, error) {
	opt["q"] = location
	opt["date_format"] = ""

	text, err := w.fetch("past-weather", opt)
	if err != nil {
		return nil, err
	}

	var o *PastLocal = new(PastLocal)
	err = xml.Unmarshal(text, o)
	if err != nil {
		return o, err
	}

	if o.Error != nil {
		return o, errors.New(*o.Error)
	}

	return o, nil
}

// Fetch historical marine weather information for location.
//
// Supported options are (defaults marked with *):
//   date     Start date (YYYY-mm-dd)
//   enddate  End date (YYYY-mm-dd)
//   tp       Number of hours in detailed forecast (1, *3, 6, 12, 24)
//   tide     Include tide information (yes, *no)
func (w *WWO) GetPastMarine(location string, opt map[string]string) (*PastMarine, error) {
	opt["q"] = location
	opt["date_format"] = ""

	text, err := w.fetch("past-marine", opt)
	if err != nil {
		return nil, err
	}

	var o *PastMarine = new(PastMarine)
	err = xml.Unmarshal(text, o)
	if err != nil {
		return o, err
	}

	if o.Error != nil {
		return o, errors.New(*o.Error)
	}

	return o, nil
}

// Look up locations.
//
// Supported options are (defaults marked with *):
//   num_of_results  Number of results to return (1-50, *10)
//   timezone        Include timezone information (yes, *no)
//   popular         Include only popular locations (yes, *no)
//   wct             Limit locations to type (ski, cricket, football, golf, fishing)
func (w *WWO) GetSearch(location string, opt map[string]string) (*Search, error) {
	opt["q"] = location
	opt["date_format"] = ""

	text, err := w.fetch("search", opt)
	if err != nil {
		return nil, err
	}

	var o *Search = new(Search)
	err = xml.Unmarshal(text, o)
	if err != nil {
		return o, err
	}

	if o.Error != nil {
		return o, errors.New(*o.Error)
	}

	return o, nil
}

// Look up time zone information for location.
//
// No supported options at the moment.
func (w *WWO) GetTimeZone(location string, opt map[string]string) (*TimeZone, error) {
	opt["q"] = location
	opt["date_format"] = ""

	text, err := w.fetch("tz", opt)
	if err != nil {
		return nil, err
	}

	var o *TimeZone = new(TimeZone)
	err = xml.Unmarshal(text, o)
	if err != nil {
		return o, err
	}

	if o.Error != nil {
		return o, errors.New(*o.Error)
	}

	return o, nil
}
