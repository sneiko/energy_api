package invoices

import (
	"context"
	"fmt"

	"energy_tk/internal/domain"
)

const queryCreate = `
		INSERT INTO invoices (
		                      user_id, 
		                      invoice_number, 
		                      from_city, 
		                      to_city, 
		                      places, 
		                      weight, 
		                      volume, 
		                      sender_is_paid, 
		                      recipient_is_paid, 
		                      delivery_date_from, 
		                      delivery_date_from_formatted, 
		                      delivery_date_to, 
		                      delivery_date_to_formatted, 
		                      sender_total_price, 
		                      recipient_total_price)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) 
		RETURNING id`

func (s *Repository) Create(ctx context.Context, invoice *domain.Invoice) (int, error) {
	invoiceResult := s.db.QueryRowContext(ctx, queryCreate,
		invoice.UserID,
		invoice.InvoiceNumber,
		invoice.FromCity,
		invoice.ToCity,
		invoice.Places,
		invoice.Weight,
		invoice.Volume,
		invoice.SenderIsPaid,
		invoice.RecipientIsPaid,
		invoice.DeliveryDateFrom,
		invoice.DeliveryDateFromFormatted,
		invoice.DeliveryDateTo,
		invoice.DeliveryDateToFormatted,
		invoice.SenderTotalPrice,
		invoice.RecipientTotalPrice)

	var invoiceId int
	errScan := invoiceResult.Scan(&invoiceId)
	if errScan != nil {
		return 0, fmt.Errorf("Invoices.GetDetailsByUserToken - last insert id: %w", errScan)
	}

	for _, state := range invoice.States {
		_, err := s.AddState(ctx, invoiceId, &state)
		if err != nil {
			return 0, fmt.Errorf("Invoices.GetDetailsByUserToken - insert states: %w", err)
		}
	}
	return invoiceId, nil
}
