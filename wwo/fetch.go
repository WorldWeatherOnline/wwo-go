package wwo

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
