package ticket

import (
	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/service"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
)

var TicketItem = new(cTicketItem)

type cTicketItem struct{}

// User GetTicketItemById doc
// @Summary      User GetTicketItemById
// @Description  Get TicketItem By Id
// @Tags         account management
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ResponseData
// @Router       /ticket/item/{id} [get]
func (c *cTicketItem) GetTicketItemById(ctx *gin.Context) {
	id := ctx.Param("id")
	ticketItem, err := service.TicketItem().GetTicketItemById(ctx, id)
	if err != nil {
		global.Logger.Error("Can not implement UserLogin interface", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid)
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, ticketItem)
}
