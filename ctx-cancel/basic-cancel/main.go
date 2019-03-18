package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://www.google.com/", nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("Context finished: %v\n", ctx.Err())
		}
	}()

	// cancel ctx, all calls with given ctx will be canceled immediately.
	cancel()
	_, err = http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Printf("request err: %v", err)
	}
	if ctx.Err() != nil {
		log.Printf("ctx err: %v", err)
	}
}
