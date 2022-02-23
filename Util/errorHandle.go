package Util

import "github.com/gin-gonic/gin"

func ErrorHTTPResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
