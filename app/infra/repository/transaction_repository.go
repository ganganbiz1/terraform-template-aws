package repository

import (
	"context"
	"errors"

	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/repository"
	"github.com/ganganbiz1/go-echo-gorm-template/app/infra"
	"gorm.io/gorm"
)

type ctxTxType struct{}

var ctxTxKey = ctxTxType{}

type Tx struct {
	db PrimaryDB
}

func NewTransaction(db PrimaryDB) repository.Transaction {
	return &Tx{
		db: db,
	}
}

func (tx *Tx) Begin(ctx context.Context) (context.Context, error) {
	t := tx.db.WithContext(ctx).Begin()
	return context.WithValue(ctx, ctxTxKey, t), nil
}

func (tx *Tx) Commit(ctx context.Context) error {
	v := ctx.Value(ctxTxKey)
	t, ok := v.(*gorm.DB)
	if !ok {
		return errors.New("commit faild")
	}
	err := t.Commit().Error
	if err != nil {
		return infra.HandleError(err)
	}

	return nil
}

func (tx *Tx) Rollback(ctx context.Context) error {
	v := ctx.Value(ctxTxKey)
	t, ok := v.(*gorm.DB)
	if !ok {
		return errors.New("rollback faild")
	}
	err := t.Rollback().Error
	if err != nil {
		return infra.HandleError(err)
	}

	return nil
}
