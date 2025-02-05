package api

import (
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

	// router.POST("/twilio-location-status", server.TwilioLocationStatusMsg)
	// router.POST("/twilio-name-status", server.TwilioNameStatusMsg)
	router.POST("/twilio-status", server.TwilioStatusMsg)
	router.POST("/twilio-receive-msg", server.TwilioReceiveMsg)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}