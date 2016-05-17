#!/bin/bash

# Reading properties
. app.properties

# Starting Kafka
$KAFKA_ROOT/bin/kafka-topics.sh --delete --zookeeper localhost:2181 --topic example
