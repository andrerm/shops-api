package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func PaginationMiddleware(defaultPage, defaultLimit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(defaultPage)))
		if err != nil || page < 1 {
			page = defaultPage
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(defaultLimit)))
		if err != nil || limit < 1 {
			limit = defaultLimit
		}

		c.Set("pagination", Pagination{Page: page, Limit: limit})
		c.Next()
	}
}
