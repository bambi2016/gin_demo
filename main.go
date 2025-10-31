package main

import (
	"github.com/gin-gonic/gin"
	"github.com/example/user-api/controller"
	"github.com/example/user-api/repository"
	"github.com/example/user-api/service"
)

func main() {
	// Initialize repository
	userRepo := repository.NewInMemoryUserRepository()

	// Initialize service
	userService := service.NewUserService(userRepo)

	// Initialize controller
	userController := controller.NewUserController(userService)

	// Initialize Gin router
	r := gin.Default()

	// API routes
	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("", userController.CreateUser)
			users.GET("/:id", userController.GetUserByID)
			users.GET("", userController.GetAllUsers)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}
	}

	// Start server
	r.Run(":8080")
}
