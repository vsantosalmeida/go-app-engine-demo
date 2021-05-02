package stream

import (
	"fmt"
	"go-app-engine-demo/config"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// KafkaProducer contains the kafka producer client
type KafkaProducer struct {
	client *kafka.Producer
}

func NewKafkaProducer() (*KafkaProducer, error) {
	cfg := &kafka.ConfigMap{
		config.BootstrapServers: config.KafkaHost,
	}

	p, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{client: p}, nil
}

func (k *KafkaProducer) Write(msg []byte, topic string) error {
	doneChan := make(chan bool)

	go func() {
		defer close(doneChan)
		for e := range k.client.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
				return

			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	km := &kafka.Message{
		Value: msg,
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
	}

	k.client.ProduceChannel() <- km

	_ = <-doneChan

	return nil
}

func (k *KafkaProducer) Close() {
	// Wait all messages to be sent or until timeout (ms)
	k.client.Flush(1000)

	k.client.Close()
}
