package user

import (
	"github.com/gin-gonic/gin"

	"github.com/onlylight29/go-ecommerce-backend-api/internal/middlewares"
)

type TicketRouter struct {
}

func (ur *TicketRouter) InitTicketRouter(router *gin.RouterGroup) {
	// Public routes
	// ticketRouterPublic := router.Group("/user")
	// {
	// 	ticketRouterPublic.GET("/item/:id", ticket.TicketItem.GetTicketItemById)
	// }

	// Private routes
	ticketRouterPrivate := router.Group("/user")
	ticketRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		ticketRouterPrivate.GET("/get_info")
	}
}
