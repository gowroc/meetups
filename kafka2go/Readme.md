# Kafka To Go

### Starting Kafka

Create `scripts/app.properties` file according to a `.sample` template:

	KAFKA_ROOT=/path/to/kafka_2.11-0.9.0.1

Starting Kafka:

	cd scripts
	# First console
	./start-zookeeper.sh
	# Second console
	./start-kafka.sh
	# Third console
	./create-example-topic.sh

### Running Go code

Starting a printer (consumer):

	go run src/cmd/printer/main.go

Starting fetching data from Twitter (producer). It requires creating `twitter/key.go` file with appropriate Twitter API information, according to a `.sample` template.

	go run src/cmd/twitter/main.go

### Taking it up a notch

Starting console producer

	cd scripts
	./start-producer.sh

Creating a topic with two partitions (requires `delete.topic.enable=true` in Kafka server.properties):

	./delete-example-topic.sh
	./create-example-topic-2.sh

### Running

Sending a message

	freeze

results in a 10-second freeze (to test multi-consumer environment).
