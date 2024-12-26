package interfaces

import (
	"context"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
)

type AuthService interface {
	Signup(ctx context.Context, user models.User) (models.User, error)
	Login(ctx context.Context, username string, password string) (models.User, error)
}
