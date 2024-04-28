package model

import (
	"time"

	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/entity"
	"gorm.io/gorm"
)

type User struct {
	ID            int
	Email         string
	Name          string
	CognitoUserID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *gorm.DeletedAt
}

func (m *User) ConverToEntity() *entity.User {
	return entity.NewUser(
		m.ID,
		m.Email,
		m.Name,
		m.CognitoUserID,
	)
}
