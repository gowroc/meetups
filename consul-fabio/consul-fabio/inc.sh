#!/usr/bin/env bash

VALUE=$(curl -s http://localhost:8500/v1/kv/value | jsawk -a 'return this[0].Value' | base64 --decode)
echo "Current value: $VALUE"
VALUE=$((VALUE+1))
echo "Incrementing value"
curl -s -X PUT -d $VALUE http://localhost:8500/v1/kv/value
echo "Incremented value: $VALUE"