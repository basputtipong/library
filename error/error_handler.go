package liberror

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors
		if len(errs) == 0 {
			return
		}

		err := errs.Last().Err
		var httpErr *HTTPError
		if errors.As(err, &httpErr) {
			c.JSON(httpErr.Code, gin.H{
				"code":    httpErr.Code,
				"message": httpErr.Message,
				"detail":  httpErr.Detail,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "internal server error",
				"detail":  err.Error(),
			})
		}
	}
}
