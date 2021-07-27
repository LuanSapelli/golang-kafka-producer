create_topics:
	@docker-compose exec kafka kafka-topics --create --topic topic-example --partitions 1 --replication-factor 1 --if-not-exists --zookeeper zookeeper:2181

delete_topics:
	@docker-compose exec kafka kafka-topics --delete --topic topic-example --zookeeper zookeeper:2181

renew_topics:
	@docker-compose exec kafka kafka-topics --delete --topic topic-example --zookeeper zookeeper:2181
	@docker-compose exec kafka kafka-topics --create --topic topic-example --partitions 1 --replication-factor 1 --if-not-exists --zookeeper zookeeper:2181

read_messages:
	@docker-compose exec kafka kafka-console-consumer --bootstrap-server kafka:9092 --topic topic-example --from-beginning

produce_message:
	@docker-compose exec kafka kafka-console-producer --topic topic-example --bootstrap-server kafka:9092

run:
	@echo "Running Application"
	@go run main.go