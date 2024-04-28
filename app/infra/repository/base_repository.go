package repository

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository interface {
	getDB(ctx context.Context) *gorm.DB
}

type baseRepository struct {
	db PrimaryDB
}

func NewBaseRepository(db PrimaryDB) BaseRepository {
	return &baseRepository{
		db: db,
	}
}

func (r *baseRepository) getDB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTxKey)
	if tx, ok := v.(*gorm.DB); ok {
		return tx
	}
	return r.db.WithContext(ctx)
}
