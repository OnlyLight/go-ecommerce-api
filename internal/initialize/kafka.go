package initialize

import (
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "kafka:29092"
	kafkaTopic = "otp-auth-topic"
)

func InitKafka() {
	global.KafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
}

// Producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		global.Logger.Error("Failed to close Kafka producer", zap.Error(err))
	}
}
