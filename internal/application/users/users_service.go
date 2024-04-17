package users

import (
	"context"

	"energy_tk/internal/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, user domain.User) (int64, error)
}

type Service struct {
	usersRepository UsersRepository
}

func New(usersRepository UsersRepository) *Service {
	return &Service{
		usersRepository: usersRepository,
	}
}

func (s *Service) Auth(ctx context.Context, token string) error {
	_, err := s.usersRepository.Create(ctx, domain.User{Token: token})
	return err
}
