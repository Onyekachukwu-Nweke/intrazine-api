package main

import (
	"fmt"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/db"
	transportHttp "github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/http"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/post"
)

func Run() error {
	fmt.Println("Starting Our Backend API")

	db, err  := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to DB")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate to DB")
		fmt.Println(err)
	}

	fmt.Println("successfully connected and pinged database")

	postService := post.NewPostService(db)

	httpHandler := transportHttp.NewHandler(postService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Piko Blog API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}