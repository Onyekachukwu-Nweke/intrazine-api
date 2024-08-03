package user

import (
	"context"
	"fmt"
	"time"
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