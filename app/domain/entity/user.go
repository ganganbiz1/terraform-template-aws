package entity

type User struct {
	ID            int
	Email         string
	Name          string
	CognitoUserID string
}

func NewUser(id int, email, name, cognitoUserID string) *User {
	return &User{
		ID:            id,
		Email:         email,
		Name:          name,
		CognitoUserID: cognitoUserID,
	}
}
