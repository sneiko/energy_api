package invoices

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"energy_tk/internal/domain"
)

const queryGetList = `
	SELECT * FROM invoices
	LEFT JOIN users u ON users.id = invoices.user_id
	WHERE u.token = $1
`

// GetListByUserToken create new user
func (s *Repository) GetListByUserToken(ctx context.Context, token string) ([]domain.Invoice, error) {
	result, err := s.db.
		NewRaw(queryGetList, token).
		Exec(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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
