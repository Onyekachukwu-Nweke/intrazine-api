package user

import (
	"context"
	"fmt"
	"time"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/utils"
)

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	PasswordHash string `json:"-"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type UserStore interface {
	CreateUser(ctx context.Context, user User) (User, error)
	CheckUserExists(ctx context.Context, username, email string) (exists bool, field string, err error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
}

type UserService struct {
	UserStore UserStore
}

func NewUserService(store UserStore) *UserService {
	return &UserService{
		UserStore: store,
	}
}

/**
 * CreateUser
 * 
 *
 */
func (s *UserService) CreateUser(ctx context.Context, user User) (User, error) {
	insertedUser, err := s.UserStore.CreateUser(ctx, user)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}

	return insertedUser, nil
}

func (s *UserService) CheckUserExists(ctx context.Context, username, email string) (exists bool, field string, err error) {
	exists, field, err = s.UserStore.CheckUserExists(ctx, username, email)
	if err != nil {
		fmt.Println(err)
		return false, "", err
	}
	return exists, field, nil
}

func (s *UserService) Login(ctx context.Context, username, password string) (User, error) {
	user, err := s.UserStore.GetUserByUsername(ctx, username)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}

	// Check the Password
	valid := utils.CheckPasswordHash(user.PasswordHash, password)
	// fmt.Println(valid)
	if !valid {
		return User{}, fmt.Errorf("invalid credentials")
	}

	return user, nil
}