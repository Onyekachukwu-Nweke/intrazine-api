package services

import (
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
)

type UserService struct {
	Repo interfaces.UserRepo
}

func NewUserService(repo interfaces.UserRepo) *UserService {
	return &UserService{Repo: repo}
}
