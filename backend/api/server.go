package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/reubenthomasjohn/location-heatmap/db/sqlc"
	"github.com/reubenthomasjohn/location-heatmap/token"
	"github.com/reubenthomasjohn/location-heatmap/util"
)

type Server struct {
	router *gin.Engine
	store *db.Store
	tokenMaker token.Maker
	config util.Config
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://www.locatetogether.net/"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}

  	router.Use(cors.New(corsConfig))

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/twilio-status", server.twilioStatusMsg)
	router.POST("/twilio-receive-msg", server.twilioReceiveMsg)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/members", server.getMembers)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}