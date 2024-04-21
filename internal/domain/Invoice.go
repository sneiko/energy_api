package domain

type Invoice struct {
	ID                        int           `json:"id" db:"id"`
	UserID                    int           `json:"user_id" db:"user_id,notnull"`
	InvoiceNumber             string        `json:"invoice_number" db:"invoice_number"`
	FromCity                  string        `json:"from_city" db:"from_city"`
	ToCity                    string        `json:"to_city" db:"to_city"`
	Places                    int           `json:"places" db:"places"`
	Weight                    int           `json:"weight" db:"weight"`
	Volume                    float64       `json:"volume" db:"volume"`
	SenderIsPaid              bool          `json:"sender_is_paid" db:"sender_is_paid"`
	RecipientIsPaid           bool          `json:"recipient_is_paid" db:"recipient_is_paid"`
	DeliveryDateFrom          int           `json:"delivery_date_from" db:"delivery_date_from"`
	DeliveryDateFromFormatted string        `json:"delivery_date_from_formatted" db:"delivery_date_from_formatted"`
	DeliveryDateTo            int           `json:"delivery_date_to" db:"delivery_date_to"`
	DeliveryDateToFormatted   string        `json:"delivery_date_to_formatted" db:"delivery_date_to_formatted"`
	SenderTotalPrice          int           `json:"sender_total_price" db:"sender_total_price"`
	RecipientTotalPrice       int           `json:"recipient_total_price" db:"recipient_total_price"`
	States                    InvoiceStates `json:"states" db:"states"`
}

type InvoiceList []Invoice
