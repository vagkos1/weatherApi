package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const apiKey = "bd5e378503939ddaee76f12ad7a97608"

func fetchWeather(city string) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
		return data
	}

	return data
}

func main() {
	startNow := time.Now()

	cities := []string{"Toronto", "London", "Paris", "Tokyo"}

	for _, city := range cities {
		data := fetchWeather(city)
		fmt.Println("The weather data for", city, "is:", data)
	}

	fmt.Println("This operation took:", time.Since(startNow))
}
