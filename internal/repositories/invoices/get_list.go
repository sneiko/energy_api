package invoices

import (
	"context"
	"database/sql"
	"fmt"

	"energy_tk/internal/domain"
)

const queryGetList = `
	SELECT * FROM invoices
	LEFT JOIN users u ON users.id = invoices.user_id
	WHERE u.token = $1
`

// Create create new user
func (s *Repository) GetListByUserToken(ctx context.Context, token string) ([]domain.Invoice, error) {
	result, err := s.db.QueryContext(ctx, queryGetList, token)
	defer result.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			return []domain.Invoice{}, nil
		}
		return nil, fmt.Errorf("User.Create - insert: %w", err)
	}

	var invoices []domain.Invoice // todo: implement right
	if err := result.Scan(&invoices); err != nil {
		return nil, fmt.Errorf("User.Create - scan: %w", err)
	}

	return invoices, nil
}
