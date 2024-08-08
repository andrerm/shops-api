package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

func Respond(c *gin.Context, status int, message string, data interface{}, err string, meta interface{}) {
	c.JSON(status, Response{
		Status:  http.StatusText(status),
		Message: message,
		Data:    data,
		Error:   err,
		Meta:    meta,
	})
}

func RespondError(c *gin.Context, status int, message string, err string) {
	Respond(c, status, message, nil, err, nil)
}

func RespondSuccess(c *gin.Context, data interface{}, meta interface{}) {
	Respond(c, http.StatusOK, "success", data, "", meta)
}
