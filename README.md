# golang-kafka-producer 🧙🏽‍♂️

- This producer reads `kafka-message.json` file and send the file content to Kafka.

### how-to-use

- Execute `docker-compose up -d` to create a Kafka broker and Zookeeper.
- Execute `make create_topics` to create a topic.
- Execute `make run` to run the application, the message will be in Kafka.
