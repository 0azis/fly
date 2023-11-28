package server

import (
	"chat/internal/routes"
	"chat/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func StartServer() {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "http://localhost:5173"
		},
	  }))

	storeInstance := store.NewStore()
	storeInstance.Open()
	defer storeInstance.Close()

	routes.InitRoutes(app, storeInstance)
	app.Run() // listen and serve on 0.0.0.0:8080
}
