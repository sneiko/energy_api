package domain

type Invoice struct {
	ID                        int           `json:"id"`
	InvoiceNumber             string        `json:"invoice_number"`
	FromCity                  string        `json:"from_city"`
	ToCity                    string        `json:"to_city"`
	Places                    int           `json:"places"`
	Weight                    int           `json:"weight"`
	Volume                    float64       `json:"volume"`
	SenderIsPaid              bool          `json:"sender_is_paid"`
	RecipientIsPaid           bool          `json:"recipient_is_paid"`
	DeliveryDateFrom          int           `json:"delivery_date_from"`
	DeliveryDateFromFormatted string        `json:"delivery_date_from_formatted"`
	DeliveryDateTo            int           `json:"delivery_date_to"`
	DeliveryDateToFormatted   string        `json:"delivery_date_to_formatted"`
	SenderTotalPrice          int           `json:"sender_total_price"`
	RecipientTotalPrice       int           `json:"recipient_total_price"`
	States                    InvoiceStates `json:"states"`
}

type InvoiceList []Invoice
