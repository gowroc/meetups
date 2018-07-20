package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/twitchtv/twirp"

	"github.com/seblw/twirp_gowroc/internal/warehouseserver"
	"github.com/seblw/twirp_gowroc/rpc/warehouse"
)

func main() {
	server := &warehouseserver.Server{}
	hooks := &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			ctx = context.WithValue(ctx, "req_recv", time.Now())
			return ctx, nil
		},
	}

	handler := warehouse.NewWarehouseServer(server, hooks)

	log.Println("Listening...")
	http.ListenAndServe(":8080", handler)
}
