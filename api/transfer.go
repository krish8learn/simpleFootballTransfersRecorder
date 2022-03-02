package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
)

type listTransferReturns struct {
	List  []DB.Transfer
	Total int64 `json:"total"`
}

type createTransfer struct {
	Season              int64  `json:"season" binding:"required"`
	PlayerName          string `json:"player_name" binding:"required"`
	SourceClubName      string `json:"source_club_name" binding:"required"`
	DestinationClubName string `json:"destination_club_name" binding:"required"`
	Amount              int64  `json:"amount" binding:"required"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	//check input
	var transferCreate createTransfer
	if bindErr := ctx.ShouldBindJSON(&transferCreate); bindErr != nil {
		//input does follow proper json
		ctx.JSON(http.StatusBadRequest, Util.ErrorHTTPResponse(bindErr))
		return
	}

	//check player exist or not
	playerDB, DBError := server.transaction.GetplayerByName(ctx, transferCreate.PlayerName)
	if DBError != nil {
		// error present , player does not exist in table
		ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(transferCreate.PlayerName+" no data found"))
		return
	}

	//source club checking
	sourceFootballClubDB, DBError := server.transaction.GetfootballclubByName(context.Background(), transferCreate.SourceClubName)
	if DBError != nil {
		// error present , club does not exist in table
		ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(transferCreate.SourceClubName+" no data found"))
		return
	}

	if playerDB.FootballclubID != sourceFootballClubDB.FcID {
		//source club and player club are not same
		ctx.JSON(http.StatusConflict, Util.ErrorHTTPCustomConflictResponse())
		return
	}

	//destination club checking
	destinationFootballclub, DBError := server.transaction.GetfootballclubByName(ctx, transferCreate.DestinationClubName)
	if DBError != nil {
		// error present , club does not exist in table
		ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(transferCreate.DestinationClubName+" no data found"))
		return
	}

	//now we perform transaction
	arg := DB.TransferTxParams{
		Season:            transferCreate.Season,
		PlayerID:          playerDB.PID,
		SourceClubID:      sourceFootballClubDB.FcID,
		DestinationClubID: destinationFootballclub.FcID,
		Amount:            transferCreate.Amount,
	}
	result, txError := server.transaction.TransferTx(ctx, arg)
	if txError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) playerNameTransfer(ctx *gin.Context) {
	playerName := ctx.Param("name")

	//check player availability
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

	//get transfer
	dbTransfer, DBError := server.transaction.GettransferByPlayerid(ctx, dbPlayer.PID)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(playerName+" no footballclub data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, dbTransfer)
}

type listTransfers struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTransfers(ctx *gin.Context) {
	var transferList listTransfers

	if bindErr := ctx.ShouldBindQuery(&transferList); bindErr != nil {
		ctx.JSON(http.StatusBadRequest, Util.ErrorHTTPResponse(bindErr))
		// fmt.Println("bind ", bindErr)
		return
	}

	// fmt.Println("---> ", transferList)

	arg := DB.GettransferListParams{
		Offset: (transferList.PageID - 1) * transferList.PageSize,
		Limit:  int32(transferList.PageSize),
	}

	// fmt.Println("--->", arg)

	dbTransferList, DBError := server.transaction.GettransferList(ctx, arg)
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

	// fmt.Println("--->", dbTransferList)
	result := &listTransferReturns{
		List:  dbTransferList,
		Total: int64(len(dbTransferList)),
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) maxTransfer(ctx *gin.Context) {
	dbMaxTansfer, DBError := server.transaction.Highesttransfer(ctx)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, dbMaxTansfer)

}

type updateTransfer struct {
	PlayerName string `json:"player_name" binding:"required"`
	Amount     int64  `json:"amount" binding:"required"`
}

func (server *Server) amountTransfer(ctx *gin.Context) {
	var transferInput updateTransfer

	//checking input from the user
	if bindErr := ctx.ShouldBindJSON(&transferInput); bindErr != nil {
		ctx.JSON(http.StatusBadRequest, Util.ErrorHTTPResponse(bindErr))
		// fmt.Println("bind ", bindErr)
		return
	}

	//need to get the t_id
	//get the player_id
	dbPlayer, DBError := server.transaction.GetplayerByName(ctx, transferInput.PlayerName)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(transferInput.PlayerName+" no data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	//get transfer_id
	arg := DB.LatesttransferParams{
		PlayerID:        dbPlayer.PID,
		DestinationClub: dbPlayer.FootballclubID,
	}
	dbTransfer, DBError := server.transaction.Latesttransfer(ctx, arg)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(transferInput.PlayerName+" no transferdata found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	argUpdate := DB.UpdatetransferParams{
		TID:    dbTransfer.TID,
		Amount: transferInput.Amount,
	}

	DBError = server.transaction.Updatetransfer(ctx, argUpdate)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, "Transfer Amount Updated")
}

func (server *Server) removeTransfer(ctx *gin.Context) {
	playerName := ctx.Param("name")

	//check the player in table
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

	//get the latest transfer
	dbTransfer, DBError := server.transaction.GetLasttransferByPlayerid(ctx, dbPlayer.PID)
	if DBError != nil {
		//error present, check type of error
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(playerName+" no transfer data found"))
			// fmt.Println("player ", DBError)
			return
		}
		//error is different
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	DBError = server.transaction.Deletetransfer(ctx, dbTransfer.TID)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, "Deleted")
}
