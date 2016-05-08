#!/bin/bash

# Reading properties
. app.properties

# Starting Kafka
$KAFKA_ROOT/bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic example
