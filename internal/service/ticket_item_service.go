package service

import (
	"context"

	"github.com/onlylight29/go-ecommerce-backend-api/internal/model"
)

type (
	ITicketName interface {
	}

	ITicketItem interface {
		GetTicketItemById(ctx context.Context, ticketId string) (out model.TicketItemsOutput, err error)
	}
)

var (
	localTicketName ITicketName
	localTicketItem ITicketItem
)

func TicketName() ITicketName {
	if localTicketName == nil {
		panic("implement localTicketName not found from Interface ITicketName")
	}

	return localTicketName
}

func InitTicketName(i ITicketName) {
	localTicketName = i
}

func TicketItem() ITicketItem {
	if localTicketItem == nil {
		panic("implement localTicketItem not found from Interface ITicketItem")
	}

	return localTicketItem
}

func InitTicketItem(i ITicketItem) {
	localTicketItem = i
}
