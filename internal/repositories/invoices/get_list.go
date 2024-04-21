package invoices

import (
	"context"
	"fmt"

	"github.com/blockloop/scan"

	"energy_tk/internal/domain"
)

const getList = `
	SELECT * FROM invoices
`

func (s *Repository) GetList(ctx context.Context) (domain.InvoiceList, error) {
	result, err := s.db.QueryContext(ctx, getList)
	if err != nil {
		return nil, fmt.Errorf("Invoices.GetList - select: %w", err)
	}
	defer result.Close()

	var invoices domain.InvoiceList
	if err := scan.Rows(&invoices, result); err != nil {
		return nil, fmt.Errorf("User.Create - scan: %w", err)
	}

	return invoices, nil
}
