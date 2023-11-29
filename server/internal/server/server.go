package server

import (
	"chat/internal/routes"
	"chat/internal/store"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-contrib/cors"
)

func StartServer() {
	app := gin.Default()

	app.Use(CORS())

	storeInstance := store.NewStore()
	storeInstance.Open()
	defer storeInstance.Close()

	routes.InitRoutes(app, storeInstance)
	app.Run() // listen and serve on 0.0.0.0:8080
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}