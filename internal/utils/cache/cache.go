package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils"
)

func GetCache(ctx context.Context, key string, obj interface{}) error {
	rs, err := global.RDB.Get(ctx, key).Result()

	if _, errMsg := utils.HandleGetKeyRedis(rs, err); errMsg != nil {
		return errMsg
	}

	if err := json.Unmarshal([]byte(rs), obj); err != nil {
		return fmt.Errorf("failed to Unmarshal")
	}

	return nil
}
