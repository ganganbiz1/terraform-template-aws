package service

import (
	"context"
	"log"

	"github.com/ganganbiz1/go-echo-gorm-template/app/domain"
	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/entity"
	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/repository"
)

type IfUserService interface {
	Create(ctx context.Context, e *entity.User) error
	Search(ctx context.Context, id int) (*entity.User, error)
}

type UserService struct {
	userRepo repository.IfUserRepository
}

func NewUserService(
	userRepo repository.IfUserRepository,
) IfUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Create(ctx context.Context, e *entity.User) error {
	if e.CognitoUserID == "" {
		// TODO cognitoに登録する処理
		log.Println("")
	}

	err := s.userRepo.Create(ctx, e)
	if err != nil {
		return domain.HandleError(domain.ErrInternal, err)
	}

	return nil
}

func (s *UserService) Search(ctx context.Context, id int) (*entity.User, error) {
	e, err := s.userRepo.Get(ctx, id)
	if err != nil && err != domain.ErrNotFound {
		return nil, domain.HandleError(domain.ErrInternal, err)
	}
	if e == nil || err == domain.ErrNotFound {
		return nil, domain.HandleError(domain.ErrNotFound, domain.ErrNotFound)
	}

	return e, nil
}
