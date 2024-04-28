package repository

import (
	"context"

	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/entity"
)

type IfUserRepository interface {
	Create(ctx context.Context, e *entity.User) error
	Get(ctx context.Context, id int) (*entity.User, error)
}
