package main

import (
	"personal-blog/common"
	"personal-blog/server"
)

// 模版加载
func init() {
	common.LoadTemplate()
}

func main() {
	server.App.Start("127.0.0.1", "8080")
}
