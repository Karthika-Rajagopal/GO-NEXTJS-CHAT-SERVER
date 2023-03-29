package router

import (
	"e/server/internal/user"
	//"e/server/internal/ws"
	//"time"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()
	r.POST("/signup", userHandler.CreateUser)
}

func Start(addr string) error {
	return r.Run(addr)
}