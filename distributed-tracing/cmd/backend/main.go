package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	zipkinsql "github.com/jcchavezs/zipkin-instrumentation-sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"

	"github.com/gowroc/meetups/distributed-tracing/tracer"
)

var httpAddr = flag.String("http", ":9090", "HTTP address")

var db *sqlx.DB

type Meetup struct {
	ID    int64
	Title string
}

func meetupsFromDB(ctx context.Context, query string) []Meetup {
	span := zipkin.SpanFromContext(ctx)
	span.Tag("query", query)
	var meetups []Meetup
	if err := db.SelectContext(ctx, &meetups, "SELECT id, title FROM meetup WHERE title ILIKE $1", "%"+query+"%"); err != nil {
		log.Fatal(err)
	}
	return meetups
}

func search() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span := zipkin.SpanFromContext(r.Context())
		span.Annotate(time.Now(), "db_start")
		meetups := meetupsFromDB(r.Context(), r.URL.Query().Get("q"))
		span.Annotate(time.Now(), "db_finished")
		for i := range meetups {
			fmt.Fprintln(w, meetups[i])
		}
	})
}

func main() {
	flag.Parse()

	trc, err := tracer.NewTracer("backend", *httpAddr)
	if err != nil {
		log.Fatal(err)
	}
	zipkinMiddleWare := zipkinhttp.NewServerMiddleware(
		trc,
		zipkinhttp.TagResponseSize(true),
	)

	driverName, err := zipkinsql.Register("postgres", trc, zipkinsql.WithAllTraceOptions())
	if err != nil {
		log.Fatal(err)
	}
	db, err = sqlx.Open(driverName, "user=pawel dbname=pawel sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", zipkinMiddleWare(search()))
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
