package auth

import (
	"context"
	"log"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/comment"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/post"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/user"
)

type AuthorizationService struct {
	postStore   post.PostStore
	commentStore comment.CommentStore
	userStore    user.UserStore
}

func NewAuthorizationService(ps post.PostStore, cs comment.CommentStore, us user.UserStore) *AuthorizationService {
	return &AuthorizationService{
		postStore:    ps,
		commentStore: cs,
		userStore:    us,
	}
}

func (a *AuthorizationService) IsUserAuthorized(ctx context.Context, userID, resourceID, resourceType string) bool {
	var ownerID string
	var err error

	switch resourceType {
	case "post":
			ownerID, err = a.postStore.GetOwnerIDByPostID(ctx, resourceID)
	case "comment":
			ownerID, err = a.commentStore.GetOwnerIDByCommentID(ctx, resourceID)
	default:
			return false
	}

	if err != nil {
			log.Printf("Authorization check failed: %v", err)
			return false
	}

	return ownerID == userID
}
