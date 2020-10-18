package middleware

import (
	"github.com/cheremisin-sergey/forum/internal/constants"
	"github.com/cheremisin-sergey/forum/internal/context"
	"github.com/gin-gonic/gin"
)

func AppContext(appContext *context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context)  {
		c.Set(constants.AppContext, appContext)
		c.Next()
	}
}
