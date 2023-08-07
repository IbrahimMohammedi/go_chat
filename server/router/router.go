package router

import (
	"server/internel/user"
	"server/internel/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
	r.POST("/ws/createroom", wsHandler.CreateRoom)
}

func Start(addr string) error {
	return r.Run(addr)
}
