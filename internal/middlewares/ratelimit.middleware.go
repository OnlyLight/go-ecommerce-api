package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/context"
	"github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)

type RateLimiter struct {
	globalRateLimiter         *limiter.Limiter
	publicAPIRateLimiter      *limiter.Limiter
	userPrivateAPIRateLimiter *limiter.Limiter
}

func NewRateLimiter() *RateLimiter {
	rateLimit := &RateLimiter{
		globalRateLimiter:         rateLimiter("100-S"),
		publicAPIRateLimiter:      rateLimiter("80-S"),
		userPrivateAPIRateLimiter: rateLimiter("50-S"),
	}

	return rateLimit
}

func rateLimiter(interval string) *limiter.Limiter {
	store, err := redisStore.NewStoreWithOptions(global.RDB, limiter.StoreOptions{
		Prefix:          "rate-limiter",
		MaxRetry:        3,
		CleanUpInterval: time.Hour,
	})

	if err != nil {
		return nil
	}

	rate, err := limiter.NewRateFromFormatted(interval)
	if err != nil {
		panic(err)
	}

	instance := limiter.New(store, rate)

	return instance
}

func (rl *RateLimiter) GlobalRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "global"
		log.Println("global-----")

		limitContext, err := rl.globalRateLimiter.Get(c, key)
		if err != nil {
			fmt.Println("failed to check rate limit GLOBAL", err)
			c.Next()
			return
		}

		if limitContext.Reached {
			log.Printf("Rate limit breached GLOBAL %s", key)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached GLOBAL, try later"})
			return
		}

		c.Next()
	}
}

func (rl *RateLimiter) PublicAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimitPath := rl.filterLimitUrlPath(urlPath)

		if rateLimitPath != nil {
			log.Println("Client IP --->", c.ClientIP())

			key := fmt.Sprintf("%s-%s", "111-222-333-444", urlPath)
			limitContext, err := rateLimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit", err)
				c.Next()
				return
			}

			if limitContext.Reached {
				log.Printf("Rate Limit breached %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached, try later"})
				return
			}
		}

		c.Next()
	}
}

func (rl *RateLimiter) UserPrivateAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimitPath := rl.filterLimitUrlPath(urlPath)

		if rateLimitPath != nil {
			userId, _ := context.GetUserIdFromUUID(c)
			key := fmt.Sprintf("%d-%s", userId, urlPath)
			limitContext, err := rateLimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit", err)
				c.Next()
				return
			}

			if limitContext.Reached {
				log.Printf("Rate Limit breached %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached, try later"})
				return
			}
		}

		c.Next()
	}
}

func (rl *RateLimiter) filterLimitUrlPath(urlPath string) *limiter.Limiter {
	if urlPath == "/v1/api/user/login" {
		return rl.publicAPIRateLimiter
	} else if urlPath == "/v1/api/user/info" {
		return rl.userPrivateAPIRateLimiter
	} else {
		return rl.globalRateLimiter
	}
}
