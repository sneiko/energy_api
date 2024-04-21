package invoices

import (
	"context"
	"fmt"

	"github.com/blockloop/scan"

	"energy_tk/internal/domain"
)

const queryGetDetails = `
	SELECT * FROM invoices WHERE id=?
`

func (s *Repository) GetDetails(ctx context.Context, id int) (*domain.Invoice, error) {
	result, err := s.db.QueryContext(ctx, queryGetDetails, id)
	if err != nil {
		return nil, fmt.Errorf("Invoices.GetDetailsByUserToken - select: %w", err)
	}
	defer result.Close()

	var invoice domain.Invoice
	if err := scan.Row(&invoice, result); err != nil {
		return nil, fmt.Errorf("Invoices.GetDetailsByUserToken - scan: %w", err)
	}

	return &invoice, nil
}
