#!/bin/bash
docker run --name some-postgres --rm -e POSTGRES_PASSWORD=postgres -p 15432:5432 -d postgres
