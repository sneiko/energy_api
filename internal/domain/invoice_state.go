package domain

type InvoiceState struct {
	ID                  int    `json:"id"`
	Title               string `json:"title"`
	MovingDate          int    `json:"moving_date"`
	MovingDateFormatted string `json:"moving_date_formatted"`
	MovingFromCity      string `json:"moving_from_city"`
	MovingToCity        string `json:"moving_to_city"`
}

type InvoiceStates []InvoiceState
