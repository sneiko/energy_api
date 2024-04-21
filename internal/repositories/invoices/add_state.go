package invoices

import (
	"context"
	"fmt"

	"energy_tk/internal/domain"
)

const queryAddState = `INSERT INTO invoice_states
	(site_id, invoice_id, title, moving_date, moving_from_city, moving_to_city) 
	VALUES(?, ?, ?, ?, ?, ?)
	RETURNING id`

func (s *Repository) AddState(ctx context.Context, invoiceID int, state *domain.InvoiceState) (int, error) {
	result := s.db.QueryRowContext(ctx, queryAddState,
		state.SiteID, invoiceID, state.Title, state.MovingDate, state.MovingFromCity, state.MovingToCity,
	)

	var createdId int
	err := result.Scan(&createdId)
	if err != nil {
		return 0, fmt.Errorf("User.Create - last insert id: %w", err)
	}
	return createdId, nil
}
