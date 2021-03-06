package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
)

type listReturns struct {
	List  []DB.Footballclub
	Total int64 `json:"total"`
}

type createFootballclub struct {
	ClubName  string `json:"club_name" binding:"required"`
	CountryFc string `json:"country_fc" binding:"required"`
	Balance   int64  `json:"balance" binding:"required"`
}

func (server *Server) createFootballclub(ctx *gin.Context) {
	var footballclubCreate createFootballclub

	//checking input from the user
	if bindErr := ctx.ShouldBindJSON(&footballclubCreate); bindErr != nil {
		ctx.JSON(http.StatusBadRequest, Util.ErrorHTTPResponse(bindErr))
		// fmt.Println("bind ", bindErr)
		return
	}

	//check whether club already exist or not
	_, DBError := server.Transaction.GetfootballclubByName(ctx, footballclubCreate.ClubName)
	if DBError != nil {
		//if the club does not exist, it must throw an error ("sql: no rows in result set" )
		// must create data
		arg := DB.CreatefootballclubParams{
			ClubName:  footballclubCreate.ClubName,
			CountryFc: footballclubCreate.CountryFc,
			Balance:   footballclubCreate.Balance,
		}

		createdFootballclub, DBError := server.Transaction.Createfootballclub(ctx, arg)
		if DBError != nil {
			ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
			// fmt.Println("player ", DBError)
			return
		}

		ctx.JSON(http.StatusOK, createdFootballclub)
		return
	}

	//no error, means data already exist in table
	ctx.JSON(http.StatusConflict, Util.ErrorHTTPCustomConflictResponse())

}

type listFootballclubs struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listFootballclubs(ctx *gin.Context) {
	var footballclubList listFootballclubs

	if bindErr := ctx.ShouldBindQuery(&footballclubList); bindErr != nil {
		ctx.JSON(http.StatusBadRequest, Util.ErrorHTTPResponse(bindErr))
		// fmt.Println("bind ", bindErr)
		return
	}

	// fmt.Println("---> ", footballclubList)

	arg := DB.ListfootballclubParams{
		Offset: (footballclubList.PageID - 1) * footballclubList.PageSize,
		Limit:  int32(footballclubList.PageSize),
	}

	// fmt.Println("--->", arg)

	dbfootballclubList, DBError := server.Transaction.Listfootballclub(ctx, arg)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse("No data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	// fmt.Println("--->", dbfootballclubList)
	result := &listReturns{
		List:  dbfootballclubList,
		Total: int64(len(dbfootballclubList)),
	}

	ctx.JSON(http.StatusOK, result)

}

func (server *Server) nameFootballclub(ctx *gin.Context) {
	//getting the value from URL path
	name := ctx.Param("name")

	dbfootballclub, DBError := server.Transaction.GetfootballclubByName(ctx, name)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(name+" no data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, dbfootballclub)
}

func (server *Server) countryFootballclubs(ctx *gin.Context) {
	//getting the value from URL path
	countryName := ctx.Param("country")

	dbfootballclub, DBError := server.Transaction.GetfootballclubByCountry(ctx, countryName)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(countryName+" no data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	result := &listReturns{
		List:  dbfootballclub,
		Total: int64(len(dbfootballclub)),
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) playerNameFootballclub(ctx *gin.Context) {
	//getting the value from URL path
	playerName := ctx.Param("player")

	//check whether the player in the DB
	dbPlayer, DBError := server.Transaction.GetplayerByName(ctx, playerName)
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

	dbfootballclub, DBError := server.Transaction.GetfootballclubByID(ctx, dbPlayer.FootballclubID)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(playerName+" footballclub data not found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, dbfootballclub)
}

type updateBalanceFootballclub struct {
	ClubName string `json:"club_name" binding:"required"`
	Balance  int64  `json:"balance" binding:"required"`
}

func (server *Server) updateBalanceFootballclub(ctx *gin.Context) {
	var footballclubUpdate updateBalanceFootballclub

	//checking input from the user
	if bindErr := ctx.ShouldBindJSON(&footballclubUpdate); bindErr != nil {
		ctx.JSON(http.StatusBadRequest, Util.ErrorHTTPResponse(bindErr))
		// fmt.Println("bind ", bindErr)
		return
	}

	//check whether club already exist or not
	existFootballClub, DBError := server.Transaction.GetfootballclubByName(ctx, footballclubUpdate.ClubName)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(footballclubUpdate.ClubName+" data not found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	arg := DB.UpdatefootballclubBalanceParams{
		FcID:    existFootballClub.FcID,
		Balance: footballclubUpdate.Balance,
	}

	DBError = server.Transaction.UpdatefootballclubBalance(ctx, arg)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, "Balance Updated")
}

func (server *Server) removeFootballclub(ctx *gin.Context) {
	//getting the value from URL path
	name := ctx.Param("name")

	dbfootballclub, DBError := server.Transaction.GetfootballclubByName(ctx, name)
	if DBError != nil {
		ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(name+" no data found"))
		return
	}

	//before performing deletion of club, need remove all the player related to that club
	DBError = server.Transaction.DeletePlayerByClubID(ctx, dbfootballclub.FcID)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	DBError = server.Transaction.Deletefootballclub(ctx, dbfootballclub.ClubName)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, "Deletion Successfull")
}
