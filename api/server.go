package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/middleware"
	"github.com/krish8learn/simpleFootballTransfersRecorder/token"
)

//server for HTTP request
type Server struct {
	transaction DB.Transaction
	tokenMaker  token.Maker
	router      *gin.Engine
	accessTime  time.Duration
}

//NewServer creates a new HTTP server and set up routing
func NewServer(transaction DB.Transaction, secureKay string, accessTime time.Duration) *Server {
	router := gin.Default()
	tokenMaker, err := token.NewJWTMaker(secureKay)
	// tokenMaker, err := token.NewPasetoMaker(secureKay)
	if err != nil {
		log.Fatalln("unable to create token maker", err)
	}

	server := &Server{
		transaction: transaction,
		tokenMaker:  tokenMaker,
		router:      router,
		accessTime:  accessTime,
	}

	authRouter := router.Group("/").Use(middleware.AuthMiddleware(server.tokenMaker))

	router.GET("/home", server.homePage)

	router.POST("/user/createUser", server.Createusers)
	router.POST("/user/login", server.Loginuser)

	authRouter.POST("/player/createPlayer", server.createPlayer)
	authRouter.GET("/player/listPlayers", server.listPlayers)
	authRouter.GET("/player/lessthanlistPlayers/:value")
	authRouter.GET("/player/higherthanlistPlayers/:value")
	authRouter.GET("/player/namePlayer/:name", server.namePlayer)
	authRouter.GET("/player/positionPlayers/:position", server.positionPlayer)
	authRouter.GET("/player/countryPlayers/:country", server.countryPlayer)
	authRouter.GET("/player/footballclubPlayers/:club", server.footballclubPlayers)
	authRouter.PUT("/player/updatevaluePlayer", server.updatePlayer)
	authRouter.DELETE("/player/removePlayer/:name", server.removePlayer)

	authRouter.POST("/footballclub/createFootballclub", server.createFootballclub)
	authRouter.GET("/footballclub/listFootballclubs", server.listFootballclubs)
	authRouter.GET("/footballclub/nameFootballclub/:name", server.nameFootballclub)
	authRouter.GET("/footballclub/countryFootballclubs/:country", server.countryFootballclubs)
	authRouter.GET("/footballclub/playernameFootballclub/:player", server.playerNameFootballclub)
	authRouter.PUT("/footballclub/updateBalanceFootballclub", server.updateBalanceFootballclub)
	authRouter.DELETE("/footballclub/removeFootballclub/:name", server.removeFootballclub)

	authRouter.POST("/transfer/createTransfer", server.createTransfer)
	authRouter.GET("/transfer/listTransfers", server.listTransfers)
	authRouter.GET("/transfer/playerNameTransfer/:name", server.playerNameTransfer)
	authRouter.GET("/transfer/maxTransfer", server.maxTransfer)
	authRouter.PUT("/transfer/amountTransfer", server.amountTransfer)
	authRouter.DELETE("/transfer/removeTransfer/:name", server.removeTransfer)

	return server
}

// to start listening of server on a port
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
