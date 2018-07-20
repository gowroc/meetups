package main

import (
	"context"
	"log"
	"net/http"

	"github.com/seblw/twirp_gowroc/rpc/warehouse"
)

func main() {
	cli := warehouse.NewWarehouseProtobufClient("http://localhost:8080", &http.Client{})

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
