package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/consul/api"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 9000, "Web service port")
	flag.Parse()

	serviceName := "web"
	ip := "127.0.0.1"
	address := fmt.Sprintf("%s:%d", ip, port)

	registerService(serviceName, address, port)
	startHttp(address)
}

// registerService registers service in consul
func registerService(serviceName, address string, port int) {
	service := &api.AgentServiceRegistration{
		ID:      serviceName + "-" + address,
		Name:    serviceName,
		Port:    port,
		Address: address,
		Tags:    []string{"http", "web"},
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + address + "/health",
			Interval: "1s",
			Timeout:  "1s",
		},
	}

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Agent().ServiceRegister(service); err != nil {
		log.Fatal(err)
	}
	log.Printf("Registered service %q in consul", serviceName)
}

// startHttp configures and starts http server
func startHttp(address string) {
	// register consul health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	// start http server
	log.Printf("Listening on %s!", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}
