#!/usr/bin/env bash

curl -s -X GET http://localhost:8500/v1/kv/hello | jsawk -a 'return this[0].Value' | base64 --decode