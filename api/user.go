package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

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
type userResp struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResp(user DB.User) userResp {
	return userResp{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
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

	createdUser, DBError := server.Transaction.Createusers(ctx, arg)
	if DBError != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}
	result := fmt.Sprintf("User Created %v, Relogin", createdUser.Username)
	ctx.JSON(http.StatusOK, result)
	return
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string   `json:"accessToken"`
	User        userResp `json:"user"`
}

func (server *Server) Loginuser(ctx *gin.Context) {
	//check input
	var userLogin loginUserRequest

	if bindErr := ctx.ShouldBindJSON(&userLogin); bindErr != nil {
		ctx.JSON(http.StatusBadRequest, Util.ErrorHTTPResponse(bindErr))
		// fmt.Println("bind ", bindErr)
		return
	}

	//check user
	user, DBError := server.Transaction.GetUsers(ctx, userLogin.Username)
	if DBError != nil {
		if DBError == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, Util.ErrorHTTPCustomNotFoundResponse(userLogin.Username+" no data found"))
			// fmt.Println("player ", DBError)
			return
		}

		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(DBError))
		// fmt.Println("player ", DBError)
		return
	}

	//user present , check the password
	passwordErr := Util.CheckPassword(userLogin.Password, user.HashedPassword)
	if passwordErr != nil {
		//wrong password inside the request
		ctx.JSON(http.StatusUnauthorized, Util.ErrorHTTPResponse(passwordErr))
		return
	}

	//password right, issue token
	tokenString, tokenErr := server.TokenMaker.CreateToken(user.Username, server.AccessTime)
	if tokenErr != nil {
		ctx.JSON(http.StatusInternalServerError, Util.ErrorHTTPResponse(tokenErr))
		return
	}

	rsp := loginUserResponse{
		AccessToken: tokenString,
		User:        newUserResp(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
