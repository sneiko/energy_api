package users

import (
	"context"
	"fmt"

	"energy_tk/internal/domain"
)

func (r *Repository) GetByToken(ctx context.Context, token string) (*domain.User, error) {
	var user *domain.User
	if err := r.db.NewSelect().
		Model(&user).
		Where("token = ?", token).
		Scan(ctx); err != nil {
		return nil, fmt.Errorf("User.GetByToken - select: %w", err)
	}

	return user, nil
}
