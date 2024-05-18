package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	db "se_school/db/sqlc"
	_ "se_school/docs"
	"se_school/util"
)

// Server serves HTTP requests for attendance service
type Server struct {
	store  db.Store
	router *gin.Engine
	dialer *util.EmailDialer
}

type ErrResponse struct {
	Error string `json:"error"`
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(store db.Store, dialer *util.EmailDialer) *Server {
	server := &Server{store: store, dialer: dialer}
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/rate", server.getCurrentRate)
	api.POST("/subscribe", server.createSubscription)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) ErrResponse {
	return ErrResponse{
		Error: err.Error(),
	}
}
