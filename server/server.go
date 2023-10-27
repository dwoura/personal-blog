package server

import (
	"log"
	"net/http"
	"personal-blog/router"
)

var App = &MyServer{}

type MyServer struct {
}

func init() {
	//路由
	router.Router()
}
func (*MyServer) Start(ip, port string) {
	//web  http协议
	server := http.Server{
		Addr: ip + ":" + port,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
