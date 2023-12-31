package views

import (
	"errors"
	"log"
	"net/http"
	"personal-blog/common"
	"personal-blog/context"
	"personal-blog/service"
	"strconv"
	"strings"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (*HTMLApi) IndexNew(ctx *context.MyContext) {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义
	//数据库查询
	if err := ctx.Request.ParseForm(); err != nil {
		log.Println("表单获取数据出错：", err)
		index.WriteError(ctx.W, errors.New("系统错误，请联系管理员!!"))
		return
	}
	//分页
	pageStr := ctx.Request.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10
	path := ctx.Request.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错：", err)
		index.WriteError(ctx.W, errors.New("系统错误，请联系管理员!!"))
	}

	index.WriteData(ctx.W, hr)
}
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义
	//数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员!!"))
		return
	}
	//分页
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员!!"))
	}

	index.WriteData(w, hr)
}
