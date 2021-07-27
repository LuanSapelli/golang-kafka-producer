package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer, err := buildProducer()
	if err != nil {
		log.Fatalf("error to build producer - %s", err)
	}

	defer producer.Close()

	go readProducerEvents(producer)

	// Set the topic name
	topic := "topic-example"

	for i := 0; i < 1; i++ {
		msg, err := os.ReadFile("./kafka-message.json")
		if err != nil {
			log.Printf("error to read file - %s", err)
		}

		publishMessage(&topic, msg, producer)
	}

	producer.Flush(15 * 1000)
}

func publishMessage(topic *string, messageBody []byte, producer *kafka.Producer) {
	partition := kafka.TopicPartition{
		Topic:     topic,
		Partition: kafka.PartitionAny,
	}

	message := kafka.Message{
		TopicPartition: partition,
		Value:          messageBody,
	}
	err := producer.Produce(&message, nil)
	if err != nil {
		log.Fatalf("error to produce message - %s", err)
	}
}

func readProducerEvents(producer *kafka.Producer) {
	for rawEvent := range producer.Events() {
		switch event := rawEvent.(type) {
		case *kafka.Message:
			if event.TopicPartition.Error != nil {
				log.Printf("error to send message - %s", event.TopicPartition.Error)
				continue
			}
			log.Println("message sent", event)
		}
	}
}

func buildProducer() (*kafka.Producer, error) {
	kafkaConfig := kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	return kafka.NewProducer(&kafkaConfig)
}
