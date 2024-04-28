package repository

import "context"

type Transaction interface {
	Begin(context.Context) (context.Context, error)
	Commit(context.Context) error
	Rollback(context.Context) error
}
