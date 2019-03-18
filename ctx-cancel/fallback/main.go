package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/nephe/talks/ctx-cancel/fallback/metrics"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dur, err := time.ParseDuration(r.Header.Get("x-sleep"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		time.Sleep(dur)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	clientA := client{
		name:  "client A",
		url:   ts.URL,
		sleep: time.Second,
	}
	clientB := client{
		name:  "client B",
		url:   ts.URL,
		sleep: 100 * time.Millisecond,
	}

	clients := []geocoder{&clientA, &clientB}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	a := address{
		Country:    "PL",
		PostalCode: "53-000",
	}

	coords, err := geocode(ctx, a, clients)
	fmt.Println(coords, err)
	metrics.Print()
}

func geocode(ctx context.Context, a address, clients []geocoder) (*coordinates, error) {
	for _, c := range clients {
		coords, err := c.Geocode(ctx, a)
		if err == nil {
			return coords, nil
		}
	}
	return nil, fmt.Errorf("failed to geocode address %v", a)
}

type geocoder interface {
	Geocode(context.Context, address) (*coordinates, error)
}

type address struct {
	Country    string
	PostalCode string
}

type coordinates struct {
	Lat float64
	Lng float64
}

type client struct {
	url   string
	name  string
	sleep time.Duration
}

func (c *client) Geocode(ctx context.Context, a address) (*coordinates, error) {
	metrics.WithLabel(c.name).Inc()
	req, err := http.NewRequest(http.MethodGet, c.url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-sleep", c.sleep.String())
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to geocode")
	}
	return &coordinates{Lat: 1, Lng: 2}, nil
}

