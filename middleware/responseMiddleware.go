package middleware

import (
	"ShopsAPI/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			utils.RespondError(c, http.StatusInternalServerError, "An internal error occurred", c.Errors.String())
			return
		}

		// If no response has been written yet, we assume it's a 404
		if !c.Writer.Written() {
			utils.RespondError(c, http.StatusNotFound, "Resource not found", "Not Found")
		}
	}
}
