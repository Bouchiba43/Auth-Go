package main

import (
	"github.com/Bouchiba43/Auth-Go/controllers"
	"github.com/Bouchiba43/Auth-Go/initializers"
	"github.com/Bouchiba43/Auth-Go/middleware"
	"github.com/Bouchiba43/Auth-Go/repositories"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()

	userController := controllers.NewUserController(repositories.NewUserRepository(initializers.DB))
	r.GET("/", userController.GetAll)
	r.POST("/signup", userController.Signup)
	r.POST("/login", userController.Login)
	r.POST("/logout", middleware.RequireAuth, userController.Logout)
	r.GET("/validate", middleware.RequireAuth, userController.Validate)

	r.Run()
}
