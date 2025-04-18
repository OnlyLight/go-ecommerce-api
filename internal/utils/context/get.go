package context

import (
	"context"
	"errors"

	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/cache"
)

type InfoUserUUID struct {
	UserId      uint64
	UserAccount string
}

func GetSubjectUUID(ctx context.Context) (string, error) {
	sUUID, ok := ctx.Value("subjectUUID").(string)
	if !ok {
		return "", errors.New("failed to get subject UUID")
	}

	return sUUID, nil
}

func GetUserIdFromUUID(ctx context.Context) (uint64, error) {
	sUUID, err := GetSubjectUUID(ctx)

	if err != nil {
		return 0, err
	}

	var infoUser InfoUserUUID
	if err := cache.GetCache(ctx, sUUID, &infoUser); err != nil {
		return 0, err
	}

	return infoUser.UserId, nil
}
