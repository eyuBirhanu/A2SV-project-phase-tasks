package repositories

import (
	"errors"
	domain "task-manager/Domain"
)

type UserRepository interface {
	Save(user *domain.User) error
	GetByEmail(email string) (*domain.User, error)
}

type InMemoryUserRepo struct {
	users []domain.User
}

func NewUserRepository() UserRepository {
	return &InMemoryUserRepo{users: []domain.User{}}
}

func (r *InMemoryUserRepo) Save(user *domain.User) error {
	r.users = append(r.users, *user)
	return nil
}

func (r *InMemoryUserRepo) GetByEmail(email string) (*domain.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}
