package domain

import "github.com/uptrace/bun"

type InvoiceState struct {
	bun.BaseModel `bun:"table:invoice_states"`

	ID                  int    `json:"id" bun:"id, pk"`
	Title               string `json:"title" bun:"title"`
	MovingDate          int    `json:"moving_date" bun:"moving_date"`
	MovingDateFormatted string `json:"moving_date_formatted" bun:"moving_date_formatted"`
	MovingFromCity      string `json:"moving_from_city" bun:"moving_from_city"`
	MovingToCity        string `json:"moving_to_city" bun:"moving_to_city"`
}

type InvoiceStates []InvoiceState
