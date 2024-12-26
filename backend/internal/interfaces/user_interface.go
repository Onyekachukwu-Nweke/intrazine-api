package interfaces

import (
	"context"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, userID string) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
}
