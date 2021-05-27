package stream

import (
	"fmt"
	"go-app-engine-demo/config"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// kafkaProducer contains the kafka producer client
type kafkaProducer struct {
	client *kafka.Producer
}

func NewKafkaProducer() (Producer, error) {
	cfg := &kafka.ConfigMap{
		config.BootstrapServers: config.KafkaHost,
	}

	p, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil, err
	}

	return &kafkaProducer{client: p}, nil
}

func (k *kafkaProducer) Write(msg []byte, topic string) error {
	deliveryChan := make(chan kafka.Event, 10000)
	km := &kafka.Message{
		Value: msg,
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
	}

	_ = k.client.Produce(km, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)

	return nil
}

func (k *kafkaProducer) Close() {
	// Wait all messages to be sent or until timeout (ms)
	k.client.Flush(1000)

	k.client.Close()
}
