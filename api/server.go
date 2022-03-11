package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/token"
)

//server for HTTP request
type Server struct {
	transaction *DB.Transaction
	tokenMaker  token.Maker
	router      *gin.Engine
	accessTime  time.Duration
}

//NewServer creates a new HTTP server and set up routing
func NewServer(transaction *DB.Transaction, secureKay string, accessTime time.Duration) *Server {
	router := gin.Default()
	// tokenMaker, err := token.NewJWTMaker(secureKay)
	tokenMaker, err := token.NewPasetoMaker(secureKay)
	if err != nil {
		log.Fatalln("unable to create token maker", err)
	}

	server := &Server{
		transaction: transaction,
		tokenMaker:  tokenMaker,
		router:      router,
		accessTime:  accessTime,
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

	router.POST("/user/createUser", server.Createusers)
	router.POST("/user/login", server.Loginuser)

	return server
}

// to start listening of server on a port
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
