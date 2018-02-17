package main

type Weather interface {
	GetCurrentConditions(city string, state string) (CurrentConditions, error)
}

type CurrentConditions struct {
	City string `json:"city"`
	State string `json:"state"`
	Temp float64 `json:"temp"`
	Wind float64 `json:"wind"`
	WindGusts string `json:"wind_gusts"`
	WindDirection string `json:"wind_direction"`
	LastUpdated string `json:"last_updated"`
}
