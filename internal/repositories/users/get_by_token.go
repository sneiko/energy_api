package users

import (
	"context"
	"fmt"

	"github.com/blockloop/scan"

	"energy_tk/internal/domain"
)

const getByTokenQuery = `select * from users where token = ? LIMIT 1`

func (r *Repository) GetByToken(ctx context.Context, token string) (*domain.User, error) {
	result, err := r.db.QueryContext(ctx, getByTokenQuery, token)
	if err != nil {
		return nil, fmt.Errorf("User.GetByToken - query select: %w", err)
	}
	defer result.Close()

	var user domain.User
	if err := scan.Row(&user, result); err != nil {
		return nil, fmt.Errorf("User.GetByToken - scan: %w", err)
	}

	return &user, nil
}
