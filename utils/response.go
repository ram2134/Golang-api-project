package utils

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, status int, payload interface{}) {
    c.JSON(status, payload)
}

func ErrorResponse(c *gin.Context, status int, message string) {
    c.JSON(status, gin.H{"error": message})
}
