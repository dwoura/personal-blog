package api

import (
	"errors"
	"net/http"
	"personal-blog/common"
	"personal-blog/dao"
	"personal-blog/models"
	"personal-blog/service"
	"personal-blog/utils"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/") //裁去前缀
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	post, err := dao.GetPostById(pId)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	common.Success(w, post) //响应返回post页面
}
func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//获取用户id（token）,判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid
	//POST代表save操作
	method := r.Method
	params := common.GetRequestJsonParam(r)
	switch method {
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
		common.Success(w, post)
	case http.MethodPut: //处理put请求的情况 即update
		//
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
		common.Success(w, post)
	}
	//params := common.GetRequestJsonParam(r)
}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)
	common.Success(w, searchResp)
}
