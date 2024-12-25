package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:9092"
	kafkaTopic = "user_topic_vip"
)

// Consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,              // 10KB
		MaxBytes:       10e6,              // 10MB
		CommitInterval: time.Second,       // 1s
		StartOffset:    kafka.FirstOffset, // it will start from the oldest message
	})
}

// Producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStockInfo(msg, typeMsg string) *StockInfo {
	return &StockInfo{
		Message: msg,
		Type:    typeMsg,
	}
}

func actionStock(c *gin.Context) {
	stockInfo := newStockInfo(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = stockInfo

	jsonBody, _ := json.Marshal(body)

	err := kafkaProducer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	})

	if err != nil {
		global.Logger.Error("Error while writing message to kafka", zap.Error(err))
		c.JSON(500, gin.H{
			"message": "Error while writing message to kafka",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Message sent to kafka",
	})
}

func RegisterConsumerATC(id int) {
	kafkaGroupName := "consumer-group-"
	kafkaGroupId := fmt.Sprintf("%s%d", kafkaGroupName, id)
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			global.Logger.Error("Error while reading message from kafka", zap.Error(err))
			fmt.Println("Error while reading message from kafka", err)
			return
		}

		global.Logger.Info("Message received from kafka")
		fmt.Printf("Message received from kafka: time: %s \nTopic: %v \nPartition: %v \nOffset: %v", string(m.Value), m.Topic, m.Partition, m.Offset)
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.GET("/action/stock", actionStock)

	RegisterConsumerATC(1)

	r.Run(":8999")
}
