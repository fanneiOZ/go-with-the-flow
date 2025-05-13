package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, SuccessResponse{Status: true, Data: &data})
}

func Error(c *gin.Context, code int, err ErrorPayload) {
	res := ErrorResponse{ErrorPayload: err}
	res.Status = false
	c.JSON(code, res)
}
