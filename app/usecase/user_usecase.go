package usecase

import (
	"context"

	"github.com/ganganbiz1/go-echo-gorm-template/app/domain"
	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/entity"
	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/repository"
	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/service"
	"github.com/ganganbiz1/go-echo-gorm-template/app/usecase/dto/input"
)

type IfUserUsecase interface {
	Signup(ctx context.Context, dto *input.User) error
	Search(ctx context.Context, id int) (*entity.User, error)
}

type UserUsecase struct {
	userService service.IfUserService
	userRepo    repository.IfUserRepository
}

func NewUserUsecase(
	userService service.IfUserService,
	userRepo repository.IfUserRepository,
) IfUserUsecase {
	return &UserUsecase{
		userService: userService,
		userRepo:    userRepo,
	}
}

func (u *UserUsecase) Signup(ctx context.Context, dto *input.User) error {
	if err := u.userService.Create(ctx, dto.ToEntity()); err != nil {
		return handleError(domain.ErrInternal, err)
	}
	return nil
}

func (u *UserUsecase) Search(ctx context.Context, id int) (*entity.User, error) {
	e, err := u.userService.Search(ctx, id)
	if err != nil {
		return nil, handleError(domain.ErrInternal, err)
	}
	return e, nil
}
