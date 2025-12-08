package usecases

import (
	"errors"
	domain "task-manager/Domain"
	infrastructure "task-manager/Infrastructure"
	repositories "task-manager/Repositories"
)

type UserUsecase interface {
	Register(user *domain.User) error
	Login(email, password string) (string, error)
}

type userUsecase struct {
	repo       repositories.UserRepository
	jwtService infrastructure.JWTService
}

func NewUserUsecase(repo repositories.UserRepository, jwt infrastructure.JWTService) UserUsecase {
	return &userUsecase{repo: repo, jwtService: jwt}
}

func (u *userUsecase) Register(user *domain.User) error {
	hashedPwd, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPwd
	return u.repo.Save(user)
}

func (u *userUsecase) Login(email, password string) (string, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if !infrastructure.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}
	return u.jwtService.GenerateToken(user.ID)
}
