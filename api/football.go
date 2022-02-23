package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
)

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

	arg := DB.CreatefootballclubParams{
		ClubName:  footballclubCreate.ClubName,
		CountryFc: footballclubCreate.CountryFc,
		Balance:   footballclubCreate.Balance,
	}

	createdFootballclub, DBError := server.transaction.Createfootballclub(ctx, arg)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, createdFootballclub)
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

	dbfootballclubList, DBError := server.transaction.Listfootballclub(ctx, arg)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	// fmt.Println("--->", dbfootballclubList)

	ctx.JSON(http.StatusOK, dbfootballclubList)

	return
}
