package usecase

import (
	"errors"
	"strconv"

	"github.com/arwansa/echo-ent/ent"
	"github.com/arwansa/echo-ent/ent/user"
	"github.com/arwansa/echo-ent/repository"
)

var ErrInvalidUserId = errors.New("invalid user id")

type UserUsecase interface {
	Create(name, email, role string) (*ent.User, error)
	GetById(userId string) (*ent.User, error)
	UpdateById(userId string, name, email, role string) (*ent.User, error)
	DeleteById(userId string) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) Create(name, email, role string) (*ent.User, error) {
	result, err := u.userRepo.Create(name, email, getUserRole(role))
	return result, err
}

func (u *userUsecase) GetById(userId string) (*ent.User, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, ErrInvalidUserId
	}

	result, err := u.userRepo.GetById(id)
	return result, err
}

func (u *userUsecase) UpdateById(userId, name, email, role string) (*ent.User, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, ErrInvalidUserId
	}

	result, err := u.userRepo.UpdateById(id, name, email, getUserRole(role))
	return result, err
}

func (u *userUsecase) DeleteById(userId string) error {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return ErrInvalidUserId
	}

	err = u.userRepo.DeleteById(id)
	return err
}

func getUserRole(role string) user.Role {
	userRole := user.RoleEmployee
	if role == user.RoleAdmin.String() {
		userRole = user.RoleAdmin
	}
	return userRole
}
