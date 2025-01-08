package middlewares

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/auth"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. Get the request url path
		uri := ctx.Request.URL.Path
		log.Println("uri request::", uri)

		jwtToken, err := auth.ExtractBearerToken(ctx)
		if !err {
			response.ErrorResponse(ctx, response.ErrCodeAuthTokenFailed)
			return
		}

		// 2. Validate jwt token by subject
		claim, errClaim := auth.VerifyTokenSubject(jwtToken)
		if errClaim != nil {
			response.ErrorResponse(ctx, response.ErrCodeAuthTokenFailed)
			return
		}
		log.Println("Claim UUID::", claim.Subject)

		// 3. Update claims to context
		c := context.WithValue(ctx.Request.Context(), "subjectUUID", claim.Subject)
		ctx.Request = ctx.Request.WithContext(c)

		ctx.Next()
	}
}
