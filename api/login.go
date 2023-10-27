package api

import (
	"personal-blog/common"
	"personal-blog/context"
	"personal-blog/service"
)

func (*Api) Login(ctx *context.MyContext) {
	//登录api
	//接收用户名密码 service执行登录操作 返回对应json数据
	//json接收方式，不能简单地用parse.form来获取了
	params := common.GetRequestJsonParam(ctx.Request)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, _ := service.Login(userName, passwd)
	//返回登录成功信息
	common.Success(ctx.W, loginRes)
}
