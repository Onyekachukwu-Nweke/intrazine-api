package services

import (
	"context"
	"fmt"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/utils"
)

type AuthService struct {
	Repo interfaces.UserRepo
}

func NewAuthService(repo interfaces.UserRepo) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Signup(ctx context.Context, user models.User) (models.User, error) {
	insertedUser, err := s.Repo.CreateUser(ctx, user)
	if err != nil {
		return models.User{}, err
	}
	return insertedUser, nil
}

func (s *UserService) Login(ctx context.Context, username, password string) (models.User, error) {
	user, err := s.Repo.GetUserByUsername(ctx, username)
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}

	// Check the Password
	valid := utils.CheckPasswordHash(user.PasswordHash, password)
	// fmt.Println(valid)
	if !valid {
		return models.User{}, fmt.Errorf("invalid credentials")
	}

	return user, nil
}
