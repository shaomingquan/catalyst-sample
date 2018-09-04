package task

import (
	"github.com/gin-gonic/gin"
)

var PrefixOfWorld = "/World"
var MethodOfWorld = "GET"

func HandlerOfWorld(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "this is world of task",
	})
}
