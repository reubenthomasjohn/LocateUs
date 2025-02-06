package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/reubenthomasjohn/location-heatmap/db/sqlc"
)

type Server struct {
	router *gin.Engine
	store *db.Store
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://google.com"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	config.AllowAllOrigins = true

  	router.Use(cors.New(config))

	router.POST("/twilio-status", server.TwilioStatusMsg)
	router.POST("/twilio-receive-msg", server.TwilioReceiveMsg)

	router.GET("/users", server.GetUsers)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}