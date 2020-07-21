package ginUtils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireParam(c *gin.Context, key string) string {
	o := c.Param(key)
	if o == "" {
		c.AbortWithStatusJSON(
			http.StatusNotAcceptable,
			gin.H{"message" : fmt.Sprintf("missing %s in url parameters", key)},
		)

		return ""
	}

	return o
}
