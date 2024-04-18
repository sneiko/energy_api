package invoices

import (
	"context"
	"fmt"

	"energy_tk/internal/domain"
)

const queryGetDetails = `
	SELECT * FROM invoices WHERE id=$1
`

func (s *Repository) GetDetails(ctx context.Context, id int) (*domain.Invoice, error) {
	result, err := s.db.QueryContext(ctx, queryGetDetails, id)
	defer result.Close()
	if err != nil {
		return nil, fmt.Errorf("Invoices.GetDetailsByUserToken - select: %w", err)
	}

	var invoice domain.Invoice // todo: implement right
	if err := result.Scan(&invoice); err != nil {
		return nil, fmt.Errorf("Invoices.GetDetailsByUserToken - scan: %w", err)
	}

	return &invoice, nil
}
