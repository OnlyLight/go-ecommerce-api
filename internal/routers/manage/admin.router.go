package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	// Public routes
	adminRouterPublic := router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// Private routes
	adminRouterPrivate := router.Group("/admin/user")
	{
		adminRouterPrivate.GET("/active_user")
	}
}
