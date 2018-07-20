package main

import (
	"context"
	"log"
	"net/http"

	"github.com/seblw/twirp_gowroc/rpc/warehouse"
)

// curl --request "POST" \
//      --location "http://localhost:8080/twirp/twirp.gowroc.warehouse.Warehouse/GetOrders" \
//      --header "Content-Type:application/json" \
//      --data '{"status": "booked"}'

func main() {
	cli := warehouse.NewWarehouseJSONClient("http://localhost:8080", &http.Client{})

	req := &warehouse.GetOrdersReq{
		Status: "created",
	}

	resp, err := cli.GetOrders(context.Background(), req)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	if resp.GetOrders() == nil {
		log.Fatal("nil response")
	}

	for i, o := range resp.GetOrders() {
		log.Printf("Order %d: %v\n", i, o)
	}
}
