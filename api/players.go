package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
)

type createPlayer struct {
	PlayerName string `json:"player_name" binding:"required"`
	Position   string `json:"position" binding:"required"`
	CountryPl  string `json:"country_pl" binding:"required"`
	Value      int64  `json:"value" binding:"required"`
	ClubName   string `json:"club_name"`
}

func (server *Server) createPlayer(ctx *gin.Context) {
	var playerCreate createPlayer
	// var DBError error
	if bindErr := ctx.ShouldBindJSON(&playerCreate); bindErr != nil {
		ctx.JSON(http.StatusBadRequest, Util.ErrorHTTPResponse(bindErr))
		fmt.Println("bind ", bindErr)
		return
	}

	//getting footballclub_id based on name
	footballclubFromreq, DBError := server.transaction.GetfootballclubByName(ctx, playerCreate.ClubName)
	if DBError != nil {
		//footballclub not found, cannot create data for the player, error --> ("sql: no rows in result set")
		ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse())
		// fmt.Println("football ", DBError)
		return
	}

	//preparing DB struct to insert data in player table
	arg := DB.CreateplayerParams{
		PlayerName:     playerCreate.PlayerName,
		Position:       playerCreate.Position,
		CountryPl:      playerCreate.CountryPl,
		Value:          playerCreate.Value,
		FootballclubID: footballclubFromreq.FcID,
	}
	//inserting player data
	playerInserted, DBError := server.transaction.Createplayer(ctx, arg)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, playerInserted)
}

type listPlayers struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPlayers(ctx *gin.Context) {
	var playerList listPlayers

	if bindErr := ctx.ShouldBindQuery(&playerList); bindErr != nil {
		ctx.JSON(http.StatusBadRequest, Util.ErrorHTTPResponse(bindErr))
		// fmt.Println("bind ", bindErr)
		return
	}

	arg := DB.GetPlayersListParams{
		Offset: (playerList.PageID - 1) * playerList.PageSize,
		Limit:  int32(playerList.PageSize),
	}

	dbPlayerList, DBError := server.transaction.GetPlayersList(ctx, arg)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, dbPlayerList)
}
