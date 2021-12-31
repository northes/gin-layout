package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/northes/gin-layout/response"
)

func PingHandler(c *gin.Context) {
	response.Success(c, gin.H{"ping": "pong"})
}
