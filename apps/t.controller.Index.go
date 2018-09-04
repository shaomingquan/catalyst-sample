package apps

import (
	"github.com/gin-gonic/gin"
)

var PrefixOfRoot = "/"
var MethodOfRoot = "GET"

func HandlerOfRoot(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "just root",
	})
}
