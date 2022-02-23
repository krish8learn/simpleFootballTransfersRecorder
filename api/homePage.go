package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type homePageInfo struct {
	Creator          string
	AppVersion       string
	GithubID         string
	PlayerAPIs       string
	TransferAPIs     string
	FootballClubAPIs string
	SwaggerLinks     string
}

func (server *Server) homePage(ctx *gin.Context) {
	homePageInfo := homePageInfo{
		Creator:          Creator,
		AppVersion:       AppVersion,
		GithubID:         GithubID,
		PlayerAPIs:       PlayerURLAPIs,
		TransferAPIs:     TransferURLAPIs,
		FootballClubAPIs: FootballClubURLAPIs,
		SwaggerLinks:     SwaggerLinks,
	}
	ctx.JSON(http.StatusOK, homePageInfo)
}
