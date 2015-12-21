// Weather information for location.
type Weather struct {
	Temp  int  `json:"temp"`
	Windy bool `json:"windy"`
}

var cityWeather = map[string]Weather{}

func setWeather(rw http.ResponseWriter, req *http.Request) {
	city := mux.Vars(req)["city"]
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
	var w Weather
	if err := json.Unmarshal(body, &w); err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	cityWeather[city] = w
}