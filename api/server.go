package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
	"github.com/krish8learn/simpleFootballTransfersRecorder/middleware"
	"github.com/krish8learn/simpleFootballTransfersRecorder/token"
)

//server for HTTP request
type Server struct {
	Transaction DB.Transaction
	TokenMaker  token.Maker
	Router      *gin.Engine
	AccessTime  time.Duration
}

//NewServer creates a new HTTP server and set up routing
func NewServer(transaction DB.Transaction, configs Util.Config) *Server {
	router := gin.Default()
	tokenMaker, err := token.NewJWTMaker(configs.SecurityKey)
	// tokenMaker, err := token.NewPasetoMaker(configs.SecurityKey)
	if err != nil {
		log.Fatalln("unable to create token maker", err)
	}

	server := &Server{
		Transaction: transaction,
		TokenMaker:  tokenMaker,
		Router:      router,
		AccessTime:  configs.AccessTime,
	}

	if configs.AuthorizationPresent == "y" {
		withOutAuth(router, *server)
		return server
	}

	authRouter := router.Group("/").Use(middleware.AuthMiddleware(server.TokenMaker))

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
	return server.Router.Run(address)
}

func withOutAuth(router *gin.Engine, server Server) {
	log.Println("Endpoints are without authorization")
	router.GET("/home", server.homePage)

	router.POST("/user/createUser", server.Createusers)
	router.POST("/user/login", server.Loginuser)

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
}
