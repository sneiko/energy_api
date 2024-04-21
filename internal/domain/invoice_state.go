package domain

type InvoiceState struct {
	ID                  int    `json:"id" db:"id"`
	SiteID              int    `json:"site_id" db:"site_id"`
	Title               string `json:"title" db:"title"`
	MovingDate          int    `json:"moving_date" db:"moving_date"`
	MovingDateFormatted string `json:"moving_date_formatted" db:"moving_date_formatted"`
	MovingFromCity      string `json:"moving_from_city" db:"moving_from_city"`
	MovingToCity        string `json:"moving_to_city" db:"moving_to_city"`
}

type InvoiceStates []InvoiceState
