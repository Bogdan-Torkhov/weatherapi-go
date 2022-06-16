package weather

import (
	"encoding/json"
	"io"
	"net/http"
)

const link = "https://api.weatherapi.com/v1/current.json"

// Example output
// {"location":{"name":"London","region":"City of London, Greater London","country":"United Kingdom","lat":51.52,"lon":-0.11,"tz_id":"Europe/London","localtime_epoch":1655375733,"localtime":"2022-06-16 11:35"},"current":{"last_updated_epoch":1655375400,"last_updated":"2022-06-16 11:30","temp_c":25.0,"temp_f":77.0,"is_day":1,"condition":{"text":"Sunny","icon":"//cdn.weatherapi.com/weather/64x64/day/113.png","code":1000},"wind_mph":6.9,"wind_kph":11.2,"wind_degree":240,"wind_dir":"WSW","pressure_mb":1022.0,"pressure_in":30.18,"precip_mm":0.0,"precip_in":0.0,"humidity":36,"cloud":0,"feelslike_c":25.4,"feelslike_f":77.8,"vis_km":10.0,"vis_miles":6.0,"uv":5.0,"gust_mph":4.7,"gust_kph":7.6}}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name            string  `json:"name"`
	Region          string  `json:"region"`
	Country         string  `json:"country"`
	Lat             float64 `json:"lat"`
	Lon             float64 `json:"lon"`
	Tz_id           string  `json:"tz_id"`
	Localtime_epoch int     `json:"localtime_epoch"`
	Localtime       string  `json:"localtime"`
}

type Current struct {
	Last_updated_epoch int       `json:"last_updated_epoch"`
	Last_updated       string    `json:"last_updated"`
	Temp_c             float64   `json:"temp_c"`
	Temp_f             float64   `json:"temp_f"`
	Is_day             int       `json:"is_day"`
	Condition          Condition `json:"condition"`
	Wind_mph           float64   `json:"wind_mph"`
	Wind_kph           float64   `json:"wind_kph"`
	Wind_degree        int       `json:"wind_degree"`
	Wind_dir           string    `json:"wind_dir"`
	Pressure_mb        float64   `json:"pressure_mb"`
	Pressure_in        float64   `json:"pressure_in"`
	Precip_mm          float64   `json:"precip_mm"`
	Precip_in          float64   `json:"precip_in"`
	Humidity           int       `json:"humidity"`
	Cloud              int       `json:"cloud"`
	Feelslike_c        float64   `json:"feelslike_c"`
	Feelslike_f        float64   `json:"feelslike_f"`
	Vis_km             float64   `json:"vis_km"`
	Vis_miles          float64   `json:"vis_miles"`
	Uv                 float64   `json:"uv"`
	Gust_mph           float64   `json:"gust_mph"`
	Gust_kph           float64   `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

func GetWeather(apiKey string, location string) (weat Weather, err error) {
	client := http.Client{}
	r, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return
	}
	if apiKey == "" || location == "" {
		panic("Empty apiKey or location!")
	}
	// Example request :: https://api.weatherapi.com/v1/current.json?key=apiKeyHere&q=Location
	url := r.URL.Query()
	url.Add("key", apiKey)
	url.Add("q", location)
	r.URL.RawQuery = url.Encode()
	answer, err := client.Do(r)
	if err != nil {
		return
	}
	weat = Weather{}
	read, _ := io.ReadAll(answer.Body)
	err = json.Unmarshal(read, &weat)
	if err != nil {
		return
	}
	return
}

type Main struct {
	Location struct {
		Country string `json:"country"`
		Name    string `json:"name"`
		Region  string `json:"region"`
	} `json:"location"`
	Current struct {
		MainWeather string  `json:"main"`
		Cloud       int     `json:"cloud"`
		Temp        float64 `json:"temp"`
		FeelsLike   float64 `json:"feels"`
		Humidity    int     `json:"humidity"`
		WindKph     float64 `json:"wind_kph"`
	} `json:"current"`
}

// Gives out the main weather (check Main struct)

func (w Weather) GetMain() (c Main) {
	c = Main{}

	// Struct Location
	c.Location.Country = w.Location.Country
	c.Location.Name = w.Location.Name
	c.Location.Region = w.Location.Region

	// Struct Current
	c.Current.MainWeather = w.Current.Condition.Text
	c.Current.Cloud = w.Current.Cloud
	c.Current.Temp = w.Current.Temp_c
	c.Current.FeelsLike = w.Current.Feelslike_c
	c.Current.Humidity = w.Current.Humidity
	c.Current.WindKph = w.Current.Wind_kph
	return
}
