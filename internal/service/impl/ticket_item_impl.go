package impl

import (
	"context"

	"github.com/onlylight29/go-ecommerce-backend-api/internal/database"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/model"
)

type sTicketItem struct {
	r *database.Queries
}

func NewTicketItemImpl(r *database.Queries) *sTicketItem {
	return &sTicketItem{
		r: r,
	}
}

func (s *sTicketItem) GetTicketItemById(ctx context.Context, ticketId string) (out model.TicketItemsOutput, err error) {
	return out, nil
}
