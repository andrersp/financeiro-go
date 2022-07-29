package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataResponse struct {
	Err     string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"sucess"`
}

func Success(c *gin.Context, statusCode int, data interface{}) {

	if statusCode == http.StatusNoContent {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	if data != nil {
		result := DataResponse{
			Success: true,
			Data:    data,
		}
		c.JSON(statusCode, result)
	}

	c.AbortWithStatus(statusCode)
}

func Error(c *gin.Context, statusCode int, err error) {

	if statusCode == http.StatusInternalServerError {
		err = errors.New("Internal error server")
	}

	result := DataResponse{
		Success: false,
		Err:     err.Error(),
	}

	c.AbortWithStatusJSON(statusCode, result)
}
