#!/bin/bash

# Reading properties
. app.properties

# Starting Kafka
$KAFKA_ROOT/bin/zookeeper-server-start.sh $KAFKA_ROOT/config/zookeeper.properties
