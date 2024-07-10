package routes

import (
	"github.com/Bouchiba43/Auth-Go/controllers"
	"github.com/Bouchiba43/Auth-Go/initializers"
	"github.com/Bouchiba43/Auth-Go/middleware"
	"github.com/Bouchiba43/Auth-Go/repositories"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {

	userRoutes := router.Group("/users")

	userController := controllers.NewUserController(repositories.NewUserRepository(initializers.DB))

	userRoutes.GET("/", userController.GetAll)
	
	//userRoutes.GET("/users/:id", userController.GetByID)
	userRoutes.GET("/validate", middleware.RequireAuth, userController.Validate)
	userRoutes.POST("/signup", userController.Signup)
	userRoutes.POST("/login", userController.Login)
	userRoutes.POST("/logout", middleware.RequireAuth, userController.Logout)
	//userRoutes.PUT("/users/:id", userController.UpdateByID)
	//userRoutes.DELETE("/users/:id", userController.DeleteByID)

}