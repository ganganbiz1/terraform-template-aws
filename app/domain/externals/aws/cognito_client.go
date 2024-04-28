package aws

import "context"

type IfCognitoClient interface {
	Signup(ctx context.Context, email string) error
	Login(ctx context.Context, email string) error
	VerifyToken(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, email string) error
	DeleteUser(ctx context.Context, email string) error
}
