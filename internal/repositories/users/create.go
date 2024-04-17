package users

import (
	"context"
	"fmt"

	"energy_tk/internal/domain"
)

const queryCreate = `
	INSERT INTO users(token, now()) VALUES($1, $2)
	RETURNING id
`

// Create create new user
func (s *Repository) Create(ctx context.Context, user domain.User) (int64, error) {
	result, err := s.db.ExecContext(ctx, queryCreate, user.Token)
	if err != nil {
		return 0, fmt.Errorf("User.Create - insert: %w", err)
	}

	createdId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("User.Create - last insert id: %w", err)
	}

	return createdId, nil
}
