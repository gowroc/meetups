#!/bin/bash

# Reading properties
. app.properties

# Starting Kafka
$KAFKA_ROOT/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic example
