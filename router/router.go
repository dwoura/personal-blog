package router

import (
	"net/http"
	"personal-blog/api"
	"personal-blog/context"
	"personal-blog/views"
)

func Router() {
	//接收到的路径统一注册到上下文全局对象处理
	http.Handle("/", context.Context)
	//路由注册(映射) 完成路径前缀树的插入、路径参数pathArgs的填充
	context.Context.Handler("/", views.HTML.IndexNew)
	context.Context.Handler("/login", views.HTML.LoginNew)
	context.Context.Handler("/c/{id}", views.HTML.CategoryNew)
	context.Context.Handler("/p/{id}", views.HTML.DetailNew)
	//context.Context.Handler("/golang", views.HTML.PigeonholeNew)
	context.Context.Handler("/writing", views.HTML.WritingNew)
	context.Context.Handler("/pigeonhole", views.HTML.PigeonholeNew)
	context.Context.Handler("/api/v1/post", api.API.OperatePost)
	context.Context.Handler("/api/v1/post/search", api.API.SearchPost)

	context.Context.Handler("/api/v1/login", api.API.Login)
	context.Context.Handler("/api/v1/qiniu/token", api.API.QiniuToken)

	//CRUD
	context.Context.Handler("/api/v1/post/{id}", api.API.OperatePost)

	//区分返回
	//1.页面 views 2.数据json 3.静态资源
	//获取分类id http://localhost:8080/c/1 /c/后面的数字
	//http.HandleFunc("/c/", views.HTML.Category)
	//http.HandleFunc("/p/", views.HTML.Detail)
	//http.HandleFunc("/login", views.HTML.Login)
	//http.HandleFunc("/writing", views.HTML.Writing)
	//http.HandleFunc("/writing/", views.HTML.Writing)
	//http.HandleFunc("/pigeonhole/", views.HTML.Pigeonhole)
	//
	//http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	//http.HandleFunc("/api/v1/login", api.API.Login)
	//http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
