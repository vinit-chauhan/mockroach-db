package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func greet() string {
	hour := time.Now().Local().Hour()
	if hour > 4 && hour < 12 {
		return "Morning"
	} else if hour >= 12 && hour < 17 {
		return "Afternoon"
	} else if hour >= 17 && hour < 21 {
		return "Evening"
	} else {
		return "Night"
	}
}

func (w Weather) PrintOutput() {
	fmt.Printf(
		"%s ,%s: %.0fC, %s\n",
		w.Location.Name,
		w.Location.Country,
		w.Current.TempC,
		w.Current.Condition.Text,
	)
}

func (w *Weather) Fetch(api_key string, city string) *Weather {
	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=10&aqi=yes&alerts=yes", api_key, city))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("Unable to read response body")
	}

	if err := json.Unmarshal(body, w); err != nil {
		panic(err)
	}

	return w
}

func Run() {
	fmt.Printf("Good %s!!!\n", greet())
	env, err := GetVars()
	if err != nil {
		panic(fmt.Sprintf("PANIC: Unable to read env file: %s", err.Error()))
	}

	var weather Weather

	weather.Fetch(env["API_KEY"], "Windsor")

	weather.PrintOutput()
}
