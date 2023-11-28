package middleware

import (
	"chat/pkg"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header["Authorization"][0] == "Bearer" {
			c.AbortWithStatusJSON(401, "HELLO")
			return
		}

		headerToken := strings.SplitN(c.Request.Header["Authorization"][0], " ", 2)[1]
		_, expiredTime, err := pkg.GetIdentity(headerToken)

		if err != nil {
			c.AbortWithStatusJSON(409, "Ошибка")
			return
		}

		if expiredTime < time.Now().Unix() {
			c.AbortWithStatusJSON(409, "Токен просрочен")
			return
		}
		c.Next()
	}
}
