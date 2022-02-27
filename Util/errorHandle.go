package Util

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	ErrorCode    int
	ErrorMessage error
}

//as error is an interface contain Error() , must be implemented by custum struct
func (r *CustomError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.ErrorCode, r.ErrorMessage)
}

func ConflictError() error {
	return &CustomError{
		ErrorCode:    409,
		ErrorMessage: errors.New("conflict"),
	}
}

func DataNotFound() error {
	return &CustomError{
		ErrorCode:    404,
		ErrorMessage: errors.New("data not found"),
	}
}

func ErrorHTTPResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func ErrorHTTPCustomConflictResponse() gin.H {
	return gin.H{"error": ConflictError().Error()}
}

func ErrorHTTPCustomNotFoundResponse() gin.H {
	return gin.H{"error": DataNotFound().Error()}
}
