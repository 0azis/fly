package routes

import (
	"chat/internal/controllers"
	"chat/internal/middleware"
	"chat/internal/store"
	"github.com/gin-gonic/gin"
)

func messageControllers(router *gin.RouterGroup, store store.InterfaceStore) {
	controller := controllers.GetMessageControllers(store)

	messages := router.Group("/message")
	messages.POST("/", middleware.Middleware(), controller.WriteMsg)
	messages.DELETE("/", middleware.Middleware(), controller.DeleteMsg)
}
