GOFILES := $(shell find . -name '*.go')

go-get:
	go get -u -v github.com/CloudCom/goose/cmd/goose
	go get -u github.com/jteeuwen/go-bindata/...

postgres-start:
	docker run -i --rm -p 5433:5432 --name sql-talk-db -e POSTGRES_PASSWORD=pass -d postgres:9.6.2

postgres-stop:
	docker stop sql-talk-db

postgres-connect:
	docker run -it --rm --link sql-talk-db:postgres postgres psql -h postgres -U postgres

query/query.go: sql/queries/*.sql
	go-bindata -o query/query.go -pkg query sql/queries


test: $(GOFILES)
	go test -v $(go list ./... | grep -v /vendor/)


benchmark: $(GOFILES)
	go test -bench .


benchmark-long: $(GOFILES)
	go test -bench . -test.benchtime=3s -test.count=3 -timeout=30m



db-up:
	goose -path sql up


