package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vivcis/simple-bank/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	//router.POST("/transfers", server.createTransfer)
	//router.GET("/transfers/:id", server.getTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse is a helper to format error message in JSON.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
