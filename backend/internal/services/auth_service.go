package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/utils"
	"log"
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

func (s *AuthService) Login(ctx context.Context, username, password string) (models.User, error) {
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

func (s *AuthService) CheckUserExists(ctx context.Context, username, email string) (exists bool, field string, err error) {
	exists, field, err = s.Repo.CheckUserExists(ctx, username, email)
	if err != nil {
		fmt.Println(err)
		return false, "", err
	}
	return exists, field, nil
}

func (s *AuthService) ForgotPassword(ctx context.Context, username string) (string, error) {
	user, err := s.Repo.GetUserByUsername(ctx, username)
	if err != nil {
		log.Print(err)
		return "", errors.New("username not found")
	}

	resetToken, err := utils.GenerateResetToken(user.ID)
	if err != nil {
		return "", err
	}

	return resetToken, nil
}

//func (a *AuthorizationService) IsUserAuthorized(ctx context.Context, userID, resourceID, resourceType string) bool {
//	var ownerID string
//	var err error
//
//	switch resourceType {
//	case "post":
//		ownerID, err = a.postStore.GetOwnerIDByPostID(ctx, resourceID)
//	case "comment":
//		ownerID, err = a.commentStore.GetOwnerIDByCommentID(ctx, resourceID)
//	default:
//		return false
//	}
//
//	if err != nil {
//		log.Printf("Authorization check failed: %v", err)
//		return false
//	}
//
//	return ownerID == userID
//}
