package initialize

import (
	"strings"
	"time"

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
	global.KafkaProducer = getKafkaWriter()
	global.Logger.Info("Connected to Kafka")
}

// Producer
func getKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
}

// Consumer
func getKafkaReader(groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          kafkaTopic,
		MinBytes:       10e3,              // 10KB
		MaxBytes:       10e6,              // 10MB
		CommitInterval: time.Second,       // 1s
		StartOffset:    kafka.FirstOffset, // it will start from the oldest message
	})
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		global.Logger.Error("Failed to close Kafka producer", zap.Error(err))
	}
}
