package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	address string
	router  *gin.Engine
}



func CreateServer(address string, router *gin.Engine) Server {
	if router == nil {
		router = gin.Default()
	}
	return Server{address: address, router: router}
}

func Run(server Server) {
	server.router.Run(server.address)
}
