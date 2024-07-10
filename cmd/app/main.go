package main

import (
	"github.com/Bouchiba43/Auth-Go/initializers"
	"github.com/Bouchiba43/Auth-Go/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run()
}
