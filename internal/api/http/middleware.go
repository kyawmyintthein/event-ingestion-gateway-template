package http

import "github.com/gin-gonic/gin"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO(student, advanced): generate or validate a request ID and add it to logs/context.
		if id := c.GetHeader("X-Request-ID"); id != "" {
			c.Header("X-Request-ID", id)
		}
		c.Next()
	}
}
