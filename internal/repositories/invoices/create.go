package invoices

import (
	"context"
	"fmt"

	"energy_tk/internal/domain"
)

const queryCreate = `
	INSERT INTO invoices (user_id, token, now()) VALUES($1, $2, $3)
	RETURN id
`

func (s *Repository) Create(ctx context.Context, token string, invoice *domain.Invoice) (int64, error) {
	result, err := s.db.ExecContext(ctx, queryCreate, token, invoice)
	if err != nil {
		return 0, fmt.Errorf("Invoices.GetDetailsByUserToken - select: %w", err)
	}

	return result.LastInsertId()
}
