package api

import (
	"errors"
	"log"
	"net/http"
	"personal-blog/common"
	"personal-blog/context"
	"personal-blog/dao"
	"personal-blog/models"
	"personal-blog/service"
	"personal-blog/utils"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(ctx *context.MyContext) {
	pIdStr := ctx.GetPathVariable("default", "id")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(ctx.W, errors.New("类型转换出错"))
		return
	}
	post, err := dao.GetPostById(pId)
	if err != nil {
		common.Error(ctx.W, errors.New("不识别此请求路径"))
		return
	}
	common.Success(ctx.W, post) //响应返回post页面
}

func (*Api) DeletePost(ctx *context.MyContext) {
	pId, err := strconv.Atoi(ctx.GetPathVariable("default", "id"))
	if err != nil {
		log.Println(err.Error())
	}
	err = service.DeletePostById(pId)
	if err != nil {
		log.Println(err.Error())
		common.Error(ctx.W, err)
	}
	common.Success(ctx.W, "删除成功")
}

func (*Api) OperatePost(ctx *context.MyContext) {
	//获取用户id（token）,判断用户是否登录
	token := ctx.Request.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(ctx.W, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid
	//POST代表save操作
	method := ctx.Request.Method
	params := common.GetRequestJsonParam(ctx.Request)
	switch method {
	case http.MethodGet:
		API.GetPost(ctx)
	case http.MethodPost: //处理post请求的情况
		cId := params["categoryId"].(string)
		//字符串需要转换成int
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64) //底层数据是float64，要转成int
		pType := int(postType)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.SavePost(post)
		//包装成功信息并返回数据给客户端
		common.Success(ctx.W, post)
	case http.MethodPut: //处理put请求的情况 即update
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidFloat := params["pid"].(float64)
		pType := int(postType)
		pid := int(pidFloat)
		post := &models.Post{
			pid,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.UpdatePost(post)
		common.Success(ctx.W, post)
	case http.MethodDelete:
		//post请求为删除
		API.DeletePost(ctx)
	}
	//params := common.GetRequestJsonParam(r)
}

func (*Api) SearchPost(ctx *context.MyContext) {
	_ = ctx.Request.ParseForm()
	condition := ctx.Request.Form.Get("val")
	if strings.TrimPrefix(condition, " ") != "" {
		searchResp := service.SearchPost(condition)
		common.Success(ctx.W, searchResp)
		return
	}
	common.Success(ctx.W, "")
}
