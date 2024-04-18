package domain

import "github.com/uptrace/bun"

type Invoice struct {
	bun.BaseModel `bun:"table:invoices"`

	ID                        int           `json:"id" bun:"id,pk"`
	UserID                    int           `json:"user_id" bun:"user_id,notnull"`
	InvoiceNumber             string        `json:"invoice_number" bun:"invoice_number"`
	FromCity                  string        `json:"from_city" bun:"from_city"`
	ToCity                    string        `json:"to_city" bun:"to_city"`
	Places                    int           `json:"places" bun:"places"`
	Weight                    int           `json:"weight" bun:"weight"`
	Volume                    float64       `json:"volume" bun:"volume"`
	SenderIsPaid              bool          `json:"sender_is_paid" bun:"sender_is_paid"`
	RecipientIsPaid           bool          `json:"recipient_is_paid" bun:"recipient_is_paid"`
	DeliveryDateFrom          int           `json:"delivery_date_from" bun:"delivery_date_from"`
	DeliveryDateFromFormatted string        `json:"delivery_date_from_formatted" bun:"delivery_date_from_formatted"`
	DeliveryDateTo            int           `json:"delivery_date_to" bun:"delivery_date_to"`
	DeliveryDateToFormatted   string        `json:"delivery_date_to_formatted" bun:"delivery_date_to_formatted"`
	SenderTotalPrice          int           `json:"sender_total_price" bun:"sender_total_price"`
	RecipientTotalPrice       int           `json:"recipient_total_price" bun:"recipient_total_price"`
	States                    InvoiceStates `json:"states" bun:"states"`
}

type InvoiceList []Invoice
