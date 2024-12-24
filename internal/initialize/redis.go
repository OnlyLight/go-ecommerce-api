package initialize

import (
	"context"
	"fmt"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var r = global.Config.Redis
var ctx = context.Background()

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,         // 10 connections on each socket of redis (each avaialble core)
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
	}

	global.Logger.Info("Connected to Redis")
	global.RDB = rdb
}