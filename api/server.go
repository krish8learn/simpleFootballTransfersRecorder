package api

import (
	"github.com/gin-gonic/gin"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
)

//server for HTTP request
type Server struct {
	transaction *DB.Transaction
	router      *gin.Engine
}

//NewServer creates a new HTTP server and set up routing
func NewServer(transaction *DB.Transaction) *Server {
	router := gin.Default()
	server := &Server{
		transaction: transaction,
		router:      router,
	}

	router.GET("/home", server.homePage)

	router.POST("/player/createPlayer", server.createPlayer)
	router.GET("/player/listPlayers", server.listPlayers)
	router.GET("/player/lessthanlistPlayers/:value")
	router.GET("/player/higherthanlistPlayers/:value")
	router.GET("/player/namePlayer")
	router.GET("/player/positionPlayers")
	router.GET("/player/countryPlayers")
	router.GET("/player/footballclubPlayers")
	router.PUT("/player/updatevaluePlayer")
	router.DELETE("/player/removePlayer")

	router.POST("/footballclub/createFootballclub", server.createFootballclub)
	router.GET("/footballclub/listFootballclubs", server.listFootballclubs)
	router.GET("/footballclub/nameFootballclub/:name",server.nameFootballclub)
	router.GET("/footballclub/countryFootballclubs/:country",server.countryFootballclubs)
	router.GET("/footballclub/playernameFootballclub/:player", server.playerNameFootballclub)
	router.PUT("/footballclub/updateBalanceFootballclub", server.updateBalanceFootballclub)
	router.DELETE("/footballclub/removeFootballclub/:name", server.removeFootballclub)

	router.POST("/transfer/createTransfer")
	router.GET("/transfer/listTransfers")
	router.GET("/transfer/playerNameTransfer")
	router.GET("/transfer/maxTransfer")
	router.PUT("/transfer/amountTransfer/:amount")
	router.DELETE("/transfer/removeTransfer")

	return server
}

// to start listening of server on a port
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
