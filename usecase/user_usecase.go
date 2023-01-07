package usecase

import (
	"context"

	"github.com/arwansa/echo-ent/ent"
	"github.com/arwansa/echo-ent/ent/user"
	"github.com/arwansa/echo-ent/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, name, email, role string) (*ent.User, error)
	GetById(ctx context.Context, userId int) (*ent.User, error)
	UpdateById(ctx context.Context, userId int, name, email, role string) (*ent.User, error)
	DeleteById(ctx context.Context, userId int) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) Create(ctx context.Context, name, email, role string) (*ent.User, error) {
	result, err := u.userRepo.Create(ctx, name, email, getUserRole(role))
	return result, err
}

func (u *userUsecase) GetById(ctx context.Context, userId int) (*ent.User, error) {
	result, err := u.userRepo.GetById(ctx, userId)
	return result, err
}

func (u *userUsecase) UpdateById(ctx context.Context, userId int, name, email, role string) (*ent.User, error) {
	result, err := u.userRepo.UpdateById(ctx, userId, name, email, getUserRole(role))
	return result, err
}

func (u *userUsecase) DeleteById(ctx context.Context, userId int) error {
	err := u.userRepo.DeleteById(ctx, userId)
	return err
}

func getUserRole(role string) user.Role {
	userRole := user.RoleEmployee
	if role == user.RoleAdmin.String() {
		userRole = user.RoleAdmin
	}
	return userRole
}
