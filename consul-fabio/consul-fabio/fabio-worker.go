package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hashicorp/consul/api"
)

func main() {
	var addr, name, prefix string
	flag.StringVar(&addr, "addr", "127.0.0.1:5000", "host:port of the service")
	flag.StringVar(&name, "name", filepath.Base(os.Args[0]), "name of the service")
	flag.StringVar(&prefix, "prefix", "", "comma-sep list of host/path prefixes to register")
	flag.Parse()

	if prefix == "" {
		flag.Usage()
		os.Exit(1)
	}

	// register prefixes
	prefixes := strings.Split(prefix, ",")
	for _, p := range prefixes {
		http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Serving %s from %s on %s\n", r.RequestURI, name, addr)
		})
	}

	// start http server
	go func() {
		log.Printf("Listening on %s serving %s", addr, prefix)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal(err)
		}
	}()

	// register consul health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})


	// get host and port as string/int
	host, portstr, err := net.SplitHostPort(addr)
	if err != nil {
		log.Fatal(err)
	}

	port, err := strconv.Atoi(portstr)
	if err != nil {
		log.Fatal(err)
	}

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	// START OMIT
	// prefixes is a slice of strings where each value
	// represents http endpoint which should be routed in Fabio
	// build urlprefix-host/path tag list
	// e.g. urlprefix-/foo, urlprefix-/bar, ...
	var tags []string
	for _, p := range prefixes {
		tags = append(tags, "urlprefix-"+p)
	}
	serviceID := name + "-" + addr
	service := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    name,
		Port:    port,
		Address: host,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + addr + "/health",
			Interval: "1s", Timeout:  "1s",
		},
	}
	if err := client.Agent().ServiceRegister(service); err != nil {
		log.Fatal(err)
	}
	log.Printf("Registered service %q in consul with tags %q", name, strings.Join(tags, ","))
	// END OMIT

	// run until we get a signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	// deregister service
	if err := client.Agent().ServiceDeregister(serviceID); err != nil {
		log.Fatal(err)
	}
	log.Printf("Deregistered service %q in consul", name)
}
