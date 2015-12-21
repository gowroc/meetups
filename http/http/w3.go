// Weather information for location.
type Weather struct {
	Temp  int  `json:"temp"`
	Windy bool `json:"windy"`
}

func weather(rw http.ResponseWriter, req *http.Request) {
	w := Weather{Temp: 5, Windy: true}
	resp, err := json.Marshal(w)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
	rw.Write(resp)
}
