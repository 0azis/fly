package routes

import (
	"chat/internal/controllers"
	"chat/internal/store"
	"github.com/gin-gonic/gin"
)

func authControllers(router *gin.RouterGroup, store store.InterfaceStore) {
	controller := controllers.GetAuthControllers(store)

	authRouter := router.Group("/auth")
	authRouter.POST("/signup", controller.SignUp)
	authRouter.POST("/signin", controller.SignIn)
}
