package global

import (
	"database/sql"

	"github.com/onlylight29/go-ecommerce-backend-api/pkg/logger"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	RDB           *redis.Client
	MDB           *gorm.DB
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
)
