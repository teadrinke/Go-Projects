package main 

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//accesing the OpenWeatherMap API Key
type apiConfigData struct {
	OpenWeatherMapApiKey string `json: "OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json: "main"`
}

//function to load api configuration from a file
func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	// Check if there was an error reading the file
	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	//unmarshal means to convert the JSON bytes to a struct
	err = json.Unmarshal(bytes, &c) // Convert the JSON bytes to a struct
	if err != nil {
        return apiConfigData{}, err
    }
	return c, nil
}

func hello(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello from Go!\n"))
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil{
		return weatherData{}, err
	}

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?APPID="+ apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err!=nil {
		return weatherData{}, err
	}
	return d, nil
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", 
	func(w http.ResponseWriter, r *http.Request){
		
		city := strings.SplitN(r.URL.Path, "/",3)[2]
		//so basically this like divides the path into 3 parts on the basis of "/" and gets the third element
		data, err := query(city)
		if err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	//create a new server and listen on port 8080
	http.ListenAndServe(":8080", nil)
	//Use: 
	//http.Error --> // to send an error response to the client
	//log.Printf --> // to log info or errors on the server console
	//fmt.Printf --> // to print info or errors to the console

}