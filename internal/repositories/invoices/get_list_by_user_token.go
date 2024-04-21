package invoices

import (
	"context"
	"fmt"

	"github.com/blockloop/scan"

	"energy_tk/internal/domain"
)

const queryGetList = `
	SELECT * FROM invoices
	LEFT JOIN users u ON users.id = invoices.user_id
	WHERE u.token = ?
`

// GetListByUserToken create new user
func (s *Repository) GetListByUserToken(ctx context.Context, token string) (domain.InvoiceList, error) {
	result, err := s.db.QueryContext(ctx, queryGetList, token)
	if err != nil {
		return nil, fmt.Errorf("User.Create - insert: %w", err)
	}
	defer result.Close()

	var invoices domain.InvoiceList
	if err := scan.Rows(&invoices, result); err != nil {
		return nil, fmt.Errorf("User.Create - scan: %w", err)
	}

	return invoices, nil
}
