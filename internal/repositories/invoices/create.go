package invoices

import (
	"context"
	"fmt"

	"energy_tk/internal/domain"
)

func (s *Repository) Create(ctx context.Context, invoice *domain.Invoice) (int64, error) {
	result, err := s.db.
		NewInsert().
		Model(invoice).
		Exec(ctx)
	if err != nil {
		return 0, fmt.Errorf("Invoices.GetDetailsByUserToken - select: %w", err)
	}
	return result.LastInsertId()
}

func (s *Repository) AddState(ctx context.Context, state *domain.InvoiceState) error {
	_, err := s.db.
		NewInsert().
		Model(state).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("User.Create - insert: %w", err)
	}
	return nil
}
