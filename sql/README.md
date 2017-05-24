Slides and code from my "Using SQL in Go, or how I learned to program without ORMs" talk.

## Running tests

1. Start dockerized postgresql using `make postgres-start`,
2. Ensure you have `goose` and `go-bindata`:`make go-get` (for some reason `goose` will not work with go1.8 and macOS, you need go1.8.1),
3. Run up migrations: `make db-up`,
4. Run tests: `make test`.

