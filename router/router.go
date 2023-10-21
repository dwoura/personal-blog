package router

import (
	"net/http"
	"personal-blog/api"
	"personal-blog/views"
)

func Router() {
	//区分返回
	//1.页面 views 2.数据json 3.静态资源
	http.HandleFunc("/", views.HTML.Index)

	//获取分类id http://localhost:8080/c/1 /c/后面的数字
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/writing/", views.HTML.Writing)
	http.HandleFunc("/pigeonhole/", views.HTML.Pigeonhole)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
