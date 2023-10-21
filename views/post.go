package views

import (
	"errors"
	"net/http"
	"personal-blog/common"
	"personal-blog/service"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail

	//获取路径参数 分割出请求id
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/") //裁去前缀
	//此时pIdStr尾部还有".htmlx"需要裁去
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)
}
