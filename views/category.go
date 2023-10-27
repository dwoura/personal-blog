package views

import (
	"personal-blog/common"
	"personal-blog/context"
	"personal-blog/service"
	"strconv"
)

func (*HTMLApi) CategoryNew(ctx *context.MyContext) {
	categoryTemplate := common.Template.Category
	//获取分类id http://localhost:8080/c/1 /c/后面的数字
	cIdStr := ctx.GetPathVariable("default", "id")
	cId, _ := strconv.Atoi(cIdStr)
	pageStr, _ := ctx.GetForm("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryResponse, err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(ctx.W, err)
		return
	}
	categoryTemplate.WriteData(ctx.W, categoryResponse)
}
