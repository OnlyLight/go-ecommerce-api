package global

import (
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/logger"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	MDB    *gorm.DB
)
