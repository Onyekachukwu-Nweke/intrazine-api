package interfaces

import (
	"context"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
)

type AuthService interface {
	Signup(ctx context.Context, user models.User) (models.User, error)
	Login(ctx context.Context, username string, password string) (models.User, error)
	CheckUserExists(ctx context.Context, username, email string) (exists bool, field string, err error)
	ForgotPassword(ctx context.Context, username string) (string, error)
}
