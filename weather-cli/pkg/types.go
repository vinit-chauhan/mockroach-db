package pkg

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Current struct {
	TempC     float64   `json:"temp_c"`
	Condition Condition `json:"condition"`
}

type Forecast struct {
	ForecastDay []struct {
		Hour []struct {
			TimeEpoch    int64     `json:"time_epoch"`
			TempC        float64   `json:"temp_c"`
			Condition    Condition `json:"condition"`
			ChanceOfRain float64   `json:"chance_of_rain"`
		} `json:"hour"`
	} `json:"forecastday"`
}

type Condition struct {
	Text string `json:"text"`
}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}
