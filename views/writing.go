package views

import (
	"net/http"
	"personal-blog/common"
	"personal-blog/service"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	//fmt.Println(wr)
	writing.WriteData(w, wr)
}
