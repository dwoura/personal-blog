package server

import (
	"log"
	"net/http"
	"personal-blog/router"
)

var App = &MyServer{}

type MyServer struct {
}

func (*MyServer) Start(ip, port string) {
	//web  http协议
	server := http.Server{
		Addr: ip + ":" + port,
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
