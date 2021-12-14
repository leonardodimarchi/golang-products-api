package server

import (
	"golang-products-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func CreateServer() Server {
	return Server{
		port: "3000",
		server: gin.Default(),
	}
}

func (myServer *Server) Run() {
	router := routes.ConfigureRoutes(myServer.server)

	log.Print("Server is running at port: ", myServer.port)
	log.Fatal(router.Run(":" + myServer.port))
}

