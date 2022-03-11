package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
)

type createUser struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (server *Server) Createusers(ctx *gin.Context) {
	var userCreate createUser

	if bindErr := ctx.ShouldBindJSON(&userCreate); bindErr != nil {
		ctx.JSON(http.StatusBadRequest, Util.ErrorHTTPResponse(bindErr))
		// fmt.Println("bind ", bindErr)
		return
	}

	//password hashing
	hashedPassword, err := Util.HashPassword(userCreate.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(err))
		// fmt.Println("player ", DBError)
		return
	}

	arg := DB.CreateusersParams{
		Username:       userCreate.Username,
		HashedPassword: hashedPassword,
		FullName:       userCreate.FullName,
		Email:          userCreate.Email,
	}

	_, DBError := server.transaction.Createusers(ctx, arg)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	ctx.JSON(http.StatusOK, "User Created , Relogin")
	return
}
