package input

import "github.com/ganganbiz1/go-echo-gorm-template/app/domain/entity"

type User struct {
	ID       int
	Email    string
	Name     string
	Password string
}

func (d *User) ToEntity() *entity.User {
	return entity.NewUser(0, d.Email, d.Name, "")
}
