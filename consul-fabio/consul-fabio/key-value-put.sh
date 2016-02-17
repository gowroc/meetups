#!/usr/bin/env bash

curl -s -X PUT -d "Hello Wroclaw it is `date +'%H:%M:%S'`" http://localhost:8500/v1/kv/hello