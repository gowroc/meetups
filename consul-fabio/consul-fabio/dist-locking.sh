#!/usr/bin/env bash

# setup value which should be increment in distributed env
curl -s -X PUT -d "1" http://localhost:8500/v1/kv/value

# increment value with distributed lock
consul lock -verbose -try=500ms lock-name-1 ./consul-fabio/inc.sh

# check value
curl -s -X GET http://localhost:8500/v1/kv/value | jsawk -a 'return this[0].Value' | base64 --decode