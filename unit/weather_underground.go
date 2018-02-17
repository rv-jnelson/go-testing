package main

import (
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
)

const baseUrl = "http://api.wunderground.com/api/%s/conditions/q/%s/%s.json"

type weatherUndergroundAPI struct {
	apiKey string
	client http.Client
}

func (w *weatherUndergroundAPI) GetCurrentConditions(city string, state string) (CurrentConditions, error) {
	url := getUrl(w.apiKey, city, state)

	r,e := w.client.Get(url)
	if e != nil {
		return CurrentConditions{}, e
	}

	defer r.Body.Close()
	bits, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return CurrentConditions{}, e
	}

	wuCC := wuCurrentConditions{}
	e = json.Unmarshal(bits, &wuCC)
	if e != nil {
		return CurrentConditions{}, e
	}

	current := CurrentConditions{
		Temp: wuCC.CurrentObservation.TempF,
		City: city,
		State: state,
		Wind: wuCC.CurrentObservation.WindMph,
		WindGusts: wuCC.CurrentObservation.WindGustMph,
		WindDirection: wuCC.CurrentObservation.WindDir,
		LastUpdated: wuCC.CurrentObservation.ObservationTime,
	}

	return current, nil
}

func NewWeatherUndergroundAPI(apiKey string) Weather {
	wug := weatherUndergroundAPI{
		apiKey: apiKey,
		client: http.Client{
			Timeout: time.Duration(30 * time.Second),
		},
	}

	return &wug
}

func getUrl(key string, city string, state string) string {
	formattedCity := strings.Replace(city, " ", "_", -1)
	return fmt.Sprintf(baseUrl, key, formattedCity, state)
}


type wuCurrentConditions struct {
	Response struct {
		Version        string `json:"version"`
		TermsofService string `json:"termsofService"`
		Features       struct {
			Conditions int `json:"conditions"`
		} `json:"features"`
	} `json:"response"`
	CurrentObservation struct {
		Image struct {
			URL   string `json:"url"`
			Title string `json:"title"`
			Link  string `json:"link"`
		} `json:"image"`
		DisplayLocation struct {
			Full           string `json:"full"`
			City           string `json:"city"`
			State          string `json:"state"`
			StateName      string `json:"state_name"`
			Country        string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Zip            string `json:"zip"`
			Magic          string `json:"magic"`
			Wmo            string `json:"wmo"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
			Elevation      string `json:"elevation"`
		} `json:"display_location"`
		ObservationLocation struct {
			Full           string `json:"full"`
			City           string `json:"city"`
			State          string `json:"state"`
			Country        string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
			Elevation      string `json:"elevation"`
		} `json:"observation_location"`
		Estimated struct {
		} `json:"estimated"`
		StationID             string  `json:"station_id"`
		ObservationTime       string  `json:"observation_time"`
		ObservationTimeRfc822 string  `json:"observation_time_rfc822"`
		ObservationEpoch      string  `json:"observation_epoch"`
		LocalTimeRfc822       string  `json:"local_time_rfc822"`
		LocalEpoch            string  `json:"local_epoch"`
		LocalTzShort          string  `json:"local_tz_short"`
		LocalTzLong           string  `json:"local_tz_long"`
		LocalTzOffset         string  `json:"local_tz_offset"`
		Weather               string  `json:"weather"`
		TemperatureString     string  `json:"temperature_string"`
		TempF                 float64 `json:"temp_f"`
		TempC                 float64 `json:"temp_c"`
		RelativeHumidity      string  `json:"relative_humidity"`
		WindString            string  `json:"wind_string"`
		WindDir               string  `json:"wind_dir"`
		WindDegrees           int     `json:"wind_degrees"`
		WindMph               float64     `json:"wind_mph"`
		WindGustMph           string  `json:"wind_gust_mph"`
		WindKph               int     `json:"wind_kph"`
		WindGustKph           string  `json:"wind_gust_kph"`
		PressureMb            string  `json:"pressure_mb"`
		PressureIn            string  `json:"pressure_in"`
		PressureTrend         string  `json:"pressure_trend"`
		DewpointString        string  `json:"dewpoint_string"`
		DewpointF             int     `json:"dewpoint_f"`
		DewpointC             int     `json:"dewpoint_c"`
		HeatIndexString       string  `json:"heat_index_string"`
		HeatIndexF            string  `json:"heat_index_f"`
		HeatIndexC            string  `json:"heat_index_c"`
		WindchillString       string  `json:"windchill_string"`
		WindchillF            string  `json:"windchill_f"`
		WindchillC            string  `json:"windchill_c"`
		FeelslikeString       string  `json:"feelslike_string"`
		FeelslikeF            string  `json:"feelslike_f"`
		FeelslikeC            string  `json:"feelslike_c"`
		VisibilityMi          string  `json:"visibility_mi"`
		VisibilityKm          string  `json:"visibility_km"`
		Solarradiation        string  `json:"solarradiation"`
		UV                    string  `json:"UV"`
		Precip1HrString       string  `json:"precip_1hr_string"`
		Precip1HrIn           string  `json:"precip_1hr_in"`
		Precip1HrMetric       string  `json:"precip_1hr_metric"`
		PrecipTodayString     string  `json:"precip_today_string"`
		PrecipTodayIn         string  `json:"precip_today_in"`
		PrecipTodayMetric     string  `json:"precip_today_metric"`
		Icon                  string  `json:"icon"`
		IconURL               string  `json:"icon_url"`
		ForecastURL           string  `json:"forecast_url"`
		HistoryURL            string  `json:"history_url"`
		ObURL                 string  `json:"ob_url"`
		Nowcast               string  `json:"nowcast"`
	} `json:"current_observation"`
}
