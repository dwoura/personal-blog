package views

import (
	"personal-blog/common"
	"personal-blog/context"
	"personal-blog/service"
)

func (*HTMLApi) WritingNew(ctx *context.MyContext) {
	writing := common.Template.Writing
	wr := service.Writing()
	//fmt.Println(wr)
	writing.WriteData(ctx.W, wr)
}
