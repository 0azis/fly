package server

import (
	"chat/internal/routes"
	"chat/internal/store"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	app := gin.Default()

	storeInstance := store.NewStore()
	storeInstance.Open()
	defer storeInstance.Close()

	routes.InitRoutes(app, storeInstance)
	app.Run() // listen and serve on 0.0.0.0:8080
}
