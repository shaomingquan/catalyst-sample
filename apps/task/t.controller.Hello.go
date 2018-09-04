package task

import (
	"github.com/gin-gonic/gin"
)

var PrefixOfHello = "/Hello"
var MethodOfHello = "GET"

func HandlerOfHello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "this is hello of task",
	})
}
