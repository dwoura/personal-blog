package views

import (
	"personal-blog/common"
	"personal-blog/context"
	"personal-blog/service"
)

func (*HTMLApi) PigeonholeNew(ctx *context.MyContext) {
	pigeonhole := common.Template.Pigeonhole
	pigeonholeRes := service.FindPostPigeonhole()
	pigeonhole.WriteData(ctx.W, pigeonholeRes)
}
