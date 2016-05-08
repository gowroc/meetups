#!/bin/bash

# Reading properties
. app.properties

# Starting Kafka
$KAFKA_ROOT/bin/kafka-server-start.sh $KAFKA_ROOT/config/server.properties
