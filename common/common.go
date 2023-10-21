package common

//面向全体的函数，归集到common里面
import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"personal-blog/config"
	"personal-blog/models"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//可能有点耗时，耗时操作扔到协程中
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

// 固定格式 请求成功封装
func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	//fmt.Println(string(resultJson))
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

// 获取请求json的参数
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	//参数放入map中
	var params map[string]interface{}
	body, _ := io.ReadAll(r.Body) //ioutil的ReadAll已被弃用，用io的就行了
	_ = json.Unmarshal(body, &params)
	return params
}

func Error(w http.ResponseWriter, err error) {
	w.Write([]byte(err.Error()))
}
