# Distributed Tracing Demo

curl -> lb -> frontends -> backends -> DB

## Running

Create DB table and populate it with some values:

```sql
CREATE TABLE meetup (
    id SERIAL,
    title character varying NOT NULL
);
INSERT INTO meetup (title) VALUES ("Basic Go Concurrency Patterns");
INSERT INTO meetup (title) VALUES ("Writing HTTP API in Go");
INSERT INTO meetup (title) VALUES ("Why Go?");
INSERT INTO meetup (title) VALUES ("Introduction to Go");
INSERT INTO meetup (title) VALUES ("Making mobile apps with GO");
INSERT INTO meetup (title) VALUES ("Vim & Go");
INSERT INTO meetup (title) VALUES ("What's new in Go 1.6");
INSERT INTO meetup (title) VALUES ("Vendoring before and after Go 1.6");
INSERT INTO meetup (title) VALUES ("Fabio - A modern consul-aware HTTP routing load-balancer");
INSERT INTO meetup (title) VALUES ("Why Go - a bit of introduction, tooling etc.");
```

Run Zipkin in Docker.

```bash
docker run -d -p 9411:9411 openzipkin/zipkin
```

Load balancer on port 8000.

```bash
go run cmd/lb/main.go
```

Frontends on ports 8081 and 8082.

```bash
go run cmd/frontend/main.go -http localhost:8081
go run cmd/frontend/main.go -http localhost:8082
```

Backends on ports 9091 and 9092.

```bash
go run cmd/backend/main.go -http localhost:9091
go run cmd/backend/main.go -http localhost:9092
```

Make some requests:

```bash
curl 'localhost:8000/?q=go'
```
