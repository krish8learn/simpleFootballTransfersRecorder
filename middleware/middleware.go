package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
	"github.com/krish8learn/simpleFootballTransfersRecorder/token"
)

const (
	authorizationHeader     = "Authorization"
	authorizationTypeBearer = "Bearer"
	authorizationPayloadKey = "authoiztionPayload"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//we will extract the header from the token
		auth := ctx.GetHeader(authorizationHeader)
		if len(auth) == 0 {
			err := fmt.Errorf("authorization header not found")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, Util.ErrorHTTPResponse(err))
			return
		}

		//authoriztion header present
		fields := strings.Fields(auth)
		if len(fields) < 2 {
			err := fmt.Errorf("authorization header format invalid")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, Util.ErrorHTTPResponse(err))
			return
		}

		//token has proper format
		if fields[0] != authorizationTypeBearer {
			err := fmt.Errorf("authorization type invalid")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, Util.ErrorHTTPResponse(err))
			return
		}

		//authorization type is Bearer
		payload, authErr := tokenMaker.VerifyToken(fields[1])
		if authErr != nil {
			err := fmt.Errorf("invalid token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, Util.ErrorHTTPResponse(err))
			return
		}

		//valid token
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()

	}
}
