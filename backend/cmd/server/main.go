package main

import (
	"fmt"
	"log"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/config"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/repositories"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/services"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/routes"
	"github.com/gin-gonic/gin"

	// "github.com/joho/godotenv"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/db"
)

func Run(cfg *config.Config) error {
	fmt.Println("Starting Our Backend API")

	database, err := db.NewDatabase(cfg)
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
	commentRepo := repositories.NewCommentRepository(database.Client)
	userRepo := repositories.NewUserRepository(database.Client)

	// Initialize services
	postService := services.NewPostService(postRepo)
	authService := services.NewAuthService(userRepo)
	commentService := services.NewCommentService(commentRepo, postService)

	// Initialize handlers
	postHandler := handlers.NewPostHandler(postService)
	authHandler := handlers.NewAuthHandler(authService)
	commentHandler := handlers.NewCommentHandler(commentService)

	// Setup router
	server := transport.NewServer(func(router *gin.Engine) {
		routes.RegisterRoutes(router, postHandler, authHandler, commentHandler)
	})

	// Start the server
	if err := server.Serve(); err != nil {
		panic(err)
	}

	return nil
}

func main() {
	fmt.Println("Intrazine API")

	// Load the config
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := Run(cfg); err != nil {
		fmt.Println(err)
	}
}
