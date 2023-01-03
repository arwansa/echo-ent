package repository

import (
	"context"

	"github.com/arwansa/echo-ent/ent"
	"github.com/arwansa/echo-ent/ent/user"
)

type UserRepository interface {
	Create(name, email string, role user.Role) (*ent.User, error)
	GetById(userId int) (*ent.User, error)
	UpdateById(userId int, name, email string, role user.Role) (*ent.User, error)
	DeleteById(userId int) error
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client: client}
}

func (r *userRepository) Create(name, email string, role user.Role) (*ent.User, error) {
	result, err := r.client.User.
		Create().
		SetName(name).
		SetEmail(email).
		SetRole(role).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) GetById(userId int) (*ent.User, error) {
	result, err := r.client.User.
		Get(context.Background(), userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) UpdateById(userId int, name, email string, role user.Role) (*ent.User, error) {
	result, err := r.client.User.
		UpdateOneID(userId).
		SetName(name).
		SetEmail(email).
		SetRole(role).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) DeleteById(userId int) error {
	err := r.client.User.
		DeleteOneID(userId).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
