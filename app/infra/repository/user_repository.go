package repository

import (
	"context"
	"errors"
	"time"

	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/entity"
	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/repository"
	"github.com/ganganbiz1/go-echo-gorm-template/app/infra"
	"github.com/ganganbiz1/go-echo-gorm-template/app/infra/repository/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	baseRepo  BaseRepository
	replicaDB ReplicaDB
}

func NewUserRepository(
	baseRepo BaseRepository,
	replicaDB ReplicaDB,
) repository.IfUserRepository {
	return &UserRepository{
		baseRepo,
		replicaDB,
	}
}

func (r *UserRepository) Create(ctx context.Context, e *entity.User) error {
	t := time.Now()
	db := r.baseRepo.getDB(ctx)
	if err := db.Model(&model.User{}).
		Create(map[string]interface{}{
			"email":           e.Email,
			"name":            e.Name,
			"cognito_user_id": e.CognitoUserID,
			"created_at":      t,
			"updated_at":      t,
		}).Error; err != nil {
		return infra.HandleError(err)
	}
	return nil
}

func (r *UserRepository) Get(ctx context.Context, id int) (*entity.User, error) {
	var m model.User
	err := r.replicaDB.WithContext(ctx).
		Where("id = ?", id).
		First(m).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, infra.HandleError(err)
	}

	return m.ConverToEntity(), nil
}
