package repository

import (
	"context"

	"github.com/arwansa/echo-ent/ent"
	"github.com/arwansa/echo-ent/ent/user"
)

type UserRepository interface {
	Create(ctx context.Context, name, email string, role user.Role) (*ent.User, error)
	GetById(ctx context.Context, userId int) (*ent.User, error)
	UpdateById(ctx context.Context, userId int, name, email string, role user.Role) (*ent.User, error)
	DeleteById(ctx context.Context, userId int) error
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client: client}
}

func (r *userRepository) Create(ctx context.Context, name, email string, role user.Role) (*ent.User, error) {
	result, err := r.client.User.
		Create().
		SetName(name).
		SetEmail(email).
		SetRole(role).
		Save(context.Background())
	return result, err
}

func (r *userRepository) GetById(ctx context.Context, userId int) (*ent.User, error) {
	result, err := r.client.User.
		Get(context.Background(), userId)
	return result, err
}

func (r *userRepository) UpdateById(ctx context.Context, userId int, name, email string, role user.Role) (*ent.User, error) {
	result, err := r.client.User.
		UpdateOneID(userId).
		SetName(name).
		SetEmail(email).
		SetRole(role).
		Save(context.Background())
	return result, err
}

func (r *userRepository) DeleteById(ctx context.Context, userId int) error {
	err := r.client.User.
		DeleteOneID(userId).
		Exec(context.Background())
	return err
}
