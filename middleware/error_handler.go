package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func ErrorHandler(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
    return func(c *gin.Context) {
        handlerFunc(c)
        if len(c.Errors) > 0 {
            c.JSON(c.Writer.Status(), gin.H{"error": c.Errors[0].Error()})
            return
        }
        if c.Writer.Status() >= http.StatusBadRequest {
            c.JSON(c.Writer.Status(), gin.H{"error": http.StatusText(c.Writer.Status())})
        }
    }
}
