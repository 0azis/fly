package routes

import (
	"chat/internal/store"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine, store store.InterfaceStore) {
	apiGroup := app.Group("/api")

	authControllers(apiGroup, store)
	chatControllers(apiGroup, store)
	messageControllers(apiGroup, store)
}
