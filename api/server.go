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
	router.GET("/player/namePlayer/:name", server.namePlayer)
	router.GET("/player/positionPlayers/:position", server.positionPlayer)
	router.GET("/player/countryPlayers/:country", server.countryPlayer)
	router.GET("/player/footballclubPlayers/:club", server.footballclubPlayers)
	router.PUT("/player/updatevaluePlayer", server.updatePlayer)
	router.DELETE("/player/removePlayer/:name", server.removePlayer)

	router.POST("/footballclub/createFootballclub", server.createFootballclub)
	router.GET("/footballclub/listFootballclubs", server.listFootballclubs)
	router.GET("/footballclub/nameFootballclub/:name", server.nameFootballclub)
	router.GET("/footballclub/countryFootballclubs/:country", server.countryFootballclubs)
	router.GET("/footballclub/playernameFootballclub/:player", server.playerNameFootballclub)
	router.PUT("/footballclub/updateBalanceFootballclub", server.updateBalanceFootballclub)
	router.DELETE("/footballclub/removeFootballclub/:name", server.removeFootballclub)

	router.POST("/transfer/createTransfer", server.createTransfer)
	router.GET("/transfer/listTransfers", server.listTransfers)
	router.GET("/transfer/playerNameTransfer/:name", server.playerNameTransfer)
	router.GET("/transfer/maxTransfer", server.maxTransfer)
	router.PUT("/transfer/amountTransfer", server.amountTransfer)
	router.DELETE("/transfer/removeTransfer/:name", server.removeTransfer)

	return server
}

// to start listening of server on a port
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
