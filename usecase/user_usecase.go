package usecase

import (
	"github.com/arwansa/echo-ent/ent"
	"github.com/arwansa/echo-ent/ent/user"
	"github.com/arwansa/echo-ent/repository"
)

type UserUsecase interface {
	Create(name, email, role string) (*ent.User, error)
	GetById(userId int) (*ent.User, error)
	UpdateById(userId int, name, email, role string) (*ent.User, error)
	DeleteById(userId int) error
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

func (u *userUsecase) GetById(userId int) (*ent.User, error) {
	result, err := u.userRepo.GetById(userId)
	return result, err
}

func (u *userUsecase) UpdateById(userId int, name, email, role string) (*ent.User, error) {
	result, err := u.userRepo.UpdateById(userId, name, email, getUserRole(role))
	return result, err
}

func (u *userUsecase) DeleteById(userId int) error {
	err := u.userRepo.DeleteById(userId)
	return err
}

func getUserRole(role string) user.Role {
	userRole := user.RoleEmployee
	if role == user.RoleAdmin.String() {
		userRole = user.RoleAdmin
	}
	return userRole
}
