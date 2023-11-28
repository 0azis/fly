package routes

import (
	"chat/internal/controllers"
	"chat/internal/middleware"
	"chat/internal/store"
	"github.com/gin-gonic/gin"
)

func chatControllers(router *gin.RouterGroup, store store.InterfaceStore) {
	controller := controllers.GetChatControllers(store)

	chat := router.Group("/chat")
	chat.POST("/", middleware.Middleware(), controller.Create)
	chat.DELETE("/", middleware.Middleware(), controller.Delete)
	chat.GET("/", middleware.Middleware(), controller.MyChats)
	chat.GET("/:chat_id", middleware.Middleware(), controller.ChatHistory)
}
