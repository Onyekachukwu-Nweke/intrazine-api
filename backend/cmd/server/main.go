package main

import (
	"fmt"
	// "log"

	// "github.com/joho/godotenv"

	//"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/auth"
	//"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/comment"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/db"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/post"
	transportHttp "github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/http"
	//"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/user"
)

func Run() error {
	fmt.Println("Starting Our Backend API")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to DB")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate to DB")
		fmt.Println(err)
	}

	fmt.Println("successfully connected and pinged database")

	//userService := user.NewUserService(db)
	postService := post.NewPostService(db)
	//commentService := comment.NewCommentService(db)
	//authService := auth.NewAuthorizationService(postService.PostStore, commentService.CommentStore, userService.UserStore)

	httpHandler := transportHttp.NewHandler(postService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Piko Blog API")

	// Load the .env file
	// err := godotenv.Load()
	// if err != nil {
	// 		log.Fatalf("Error loading .env file")
	// }

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
