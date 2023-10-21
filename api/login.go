package api

import (
	"net/http"
	"personal-blog/common"
	"personal-blog/service"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	//登录api
	//接收用户名密码 service执行登录操作 返回对应json数据
	//json接收方式，不能简单地用parse.form来获取了
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, _ := service.Login(userName, passwd)
	//返回登录成功信息
	common.Success(w, loginRes)
}
