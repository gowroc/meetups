// Weather information for location.
type Weather struct {
	Temp  int  `json:"temp"`
	Windy bool `json:"windy"`
}

var cityWeather = map[string]Weather{}

func weather(rw http.ResponseWriter, req *http.Request) {
	city := mux.Vars(req)["city"]
	resp, err := json.Marshal(cityWeather[city])
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
	rw.Write(resp)
}
