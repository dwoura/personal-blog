package views

import (
	"errors"
	"personal-blog/common"
	"personal-blog/context"
	"personal-blog/service"
	"strconv"
)

func (*HTMLApi) DetailNew(ctx *context.MyContext) {
	detail := common.Template.Detail

	pIdStr := ctx.GetPathVariable("html", "id")

	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(ctx.W, errors.New("不识别此请求路径"))
		return
	}
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(ctx.W, errors.New("查询出错"))
		return
	}
	detail.WriteData(ctx.W, postRes)
}
