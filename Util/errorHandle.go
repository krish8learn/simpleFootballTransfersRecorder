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

func DataNotFound(input string) error {
	return &CustomError{
		ErrorCode:    404,
		ErrorMessage: errors.New(input),
	}
}

func ErrorHTTPResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func ErrorHTTPCustomConflictResponse() gin.H {
	return gin.H{"error": ConflictError().Error()}
}

func ErrorHTTPCustomNotFoundResponse(input string) gin.H {
	return gin.H{"error": DataNotFound(input).Error()}
}
