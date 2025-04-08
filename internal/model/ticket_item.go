package model

// VO: Value Object GetTicketItems return
type TicketItemsOutput struct {
	TicketId       string `json:"ticket_id"`
	TicketName     string `json:"ticket_name"`
	StockAvailable int    `json:"stock_available"`
	StockInitial   int    `json:"stock_initial"`
}

// DTO
type TicketItemsInput struct {
	TicketId string `json:"ticket_id"`
}
