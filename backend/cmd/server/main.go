package main

import (
	"fmt"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/repositories"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/services"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/routes"
	"github.com/gorilla/mux"

	// "log"

	// "github.com/joho/godotenv"

	//"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/auth"
	//"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/comment"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/db"
	//"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/post"
	//transportHttp "github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/http"
	//"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/user"
)

func Run() error {
	fmt.Println("Starting Our Backend API")

	database, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to DB")
		return err
	}
	if err := database.MigrateDB(); err != nil {
		fmt.Println("failed to migrate to DB")
		fmt.Println(err)
	}

	fmt.Println("successfully connected and pinged database")

	// Initialize repositories
	postRepo := repositories.NewPostRepository(database.Client)

	// Initialize services
	postService := services.NewPostService(postRepo)

	// Initialize handlers
	postHandler := handlers.NewPostHandler(postService)

	// Setup router
	r := mux.NewRouter()
	routes.RegisterRoutes(r, postHandler)

	//if err := httpHandler.Serve(); err != nil {
	//	return err
	//}

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
