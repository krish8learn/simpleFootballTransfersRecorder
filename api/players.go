package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
)

type playerlistReturns struct {
	List  []DB.Player
	Total int64 `json:"total"`
}

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
		// fmt.Println("bind ", bindErr)
		return
	}

	//getting footballclub_id based on name
	footballclubFromreq, DBError := server.transaction.GetfootballclubByName(ctx, playerCreate.ClubName)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(playerCreate.ClubName+" data not found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
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
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse("no data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, dbPlayerList)
}

func (server *Server) namePlayer(ctx *gin.Context) {
	playerName := ctx.Param("name")

	dbPlayer, DBError := server.transaction.GetplayerByName(ctx, playerName)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(playerName+" no data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, dbPlayer)
}

func (server *Server) positionPlayer(ctx *gin.Context) {
	position := ctx.Param("position")

	dbPlayers, DBError := server.transaction.GetplayerByPosition(ctx, position)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(position+" no data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, playerlistReturns{dbPlayers, int64(len(dbPlayers))})
}

func (server *Server) countryPlayer(ctx *gin.Context) {
	country := ctx.Param("country")

	dbPlayers, DBError := server.transaction.GetplayerByCountry(ctx, country)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(country+" no data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, playerlistReturns{dbPlayers, int64(len(dbPlayers))})
}

func (server *Server) footballclubPlayers(ctx *gin.Context) {
	club := ctx.Param("club")

	//get football club id
	dbFootballclub, DBError := server.transaction.GetfootballclubByName(ctx, club)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(club+" no data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	dbPlayers, DBError := server.transaction.GetplayerByFootballclub(ctx, dbFootballclub.FcID)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(club+" no player data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, playerlistReturns{dbPlayers, int64(len(dbPlayers))})
}

type updateValuePlayer struct {
	PlayerName       string `json:"player_name" binding:"required"`
	Value            int64  `json:"value" binding:"required"`
	FootballclubName string `json:"footballclub_name" binding:"required"`
}

func (server *Server) updatePlayer(ctx *gin.Context) {
	var playerValueUpdate updateValuePlayer

	//checking input from the user
	if bindErr := ctx.ShouldBindJSON(&playerValueUpdate); bindErr != nil {
		ctx.JSON(http.StatusBadRequest, Util.ErrorHTTPResponse(bindErr))
		// fmt.Println("bind ", bindErr)
		return
	}

	existPlayer, DBError := server.transaction.GetplayerByName(ctx, playerValueUpdate.PlayerName)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(playerValueUpdate.PlayerName+" no data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	//find new transfer club by name
	existFootballClub, DBError := server.transaction.GetfootballclubByName(ctx, playerValueUpdate.FootballclubName)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(playerValueUpdate.FootballclubName+" no data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	arg := DB.UpdateplayerParams{
		PID:            existPlayer.PID,
		Value:          playerValueUpdate.Value,
		FootballclubID: existFootballClub.FcID,
	}

	DBError = server.transaction.Updateplayer(ctx, arg)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, "Player Updated")
}

func (server *Server) removePlayer(ctx *gin.Context) {
	playerName := ctx.Param("name")

	//check whether the player exists or not
	dbPlayer, DBError := server.transaction.GetplayerByName(ctx, playerName)
	if DBError != nil {
		ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(playerName+" no data found"))
		// fmt.Println("player ", DBError)
		return
	}

	DBError = server.transaction.Deleteplayer(ctx, dbPlayer.PlayerName)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, "Deletion Successfull")
}
