package imports

import "github.com/gin-gonic/gin"
import core "github.com/shaomingquan/catalyst/core"

import "github.com/shaomingquan/catalyst-sample/apps"

import middwares "github.com/shaomingquan/catalyst-sample/middwares"

func Start_(app *core.App) {

	app.MidWare(
		"/",
		middwares.Demo(`root`),
	)

	app.Router(
		"/",
		apps.MethodOfRoot,
		apps.PrefixOfRoot,

		apps.HandlerOfRoot,
		func(ctx *gin.Context) { ctx.Next() },
	)

}

// auto generate by _, dont modify
