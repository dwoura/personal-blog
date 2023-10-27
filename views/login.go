package views

import (
	"net/http"
	"personal-blog/common"
	"personal-blog/config"
	"personal-blog/context"
)

func (*HTMLApi) LoginNew(ctx *context.MyContext) {
	login := common.Template.Login
	login.WriteData(ctx.W, config.Cfg.Viewer)
}

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login
	login.WriteData(w, config.Cfg.Viewer)
}
