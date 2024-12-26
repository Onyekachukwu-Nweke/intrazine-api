package services

import (
	"context"
	"fmt"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
)

type UserService struct {
	Repo interfaces.UserRepo
}

func NewUserService(repo interfaces.UserRepo) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CheckUserExists(ctx context.Context, username, email string) (exists bool, field string, err error) {
	exists, field, err = s.Repo.CheckUserExists(ctx, username, email)
	if err != nil {
		fmt.Println(err)
		return false, "", err
	}
	return exists, field, nil
}
