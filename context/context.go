package context

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var Context = NewContext()

// 上下文结构体
// 用于持有请求、写入、路由、路径参数
type MyContext struct {
	Request  *http.Request
	W        http.ResponseWriter
	routers  map[string]func(ctx *MyContext)
	pathArgs map[string]map[string]string
}

// 实例化
func NewContext() *MyContext {
	ctx := &MyContext{}
	ctx.routers = make(map[string]func(ctx2 *MyContext))
	ctx.pathArgs = make(map[string]map[string]string)
	return ctx
}

var UrlTree = NewTrie()

// 前缀树结构 用于路径参数匹配
type Trie struct {
	next   map[string]*Trie
	isWord bool
}

// Trie实例构造函数
func NewTrie() Trie {
	root := new(Trie)
	root.next = make(map[string]*Trie)
	root.isWord = false
	return *root
}

// 插入数据，路由根据“/”进行拆分
func (t *Trie) Insert(word string) {
	for _, v := range strings.Split(word, "/") {
		//若前缀树next[index]节点为空，则建立新节点并且接入
		if t.next[v] == nil {
			node := new(Trie)
			node.next = make(map[string]*Trie)
			node.isWord = false
			t.next[v] = node
		}
		// * 匹配所有
		// {X} 匹配路由参数 X
		if v == "*" || strings.Index(v, "{") != -1 {
			t.isWord = true
		}
		t = t.next[v]
	}
	t.isWord = true
}

// 匹配路由
func (t *Trie) Search(word string) (isHave bool, arg map[string]string) {
	arg = make(map[string]string)
	isHave = false
	//遍历分割好的字符串。例如/c/p/{id}则分出 c p {id}三个字符串
	for _, v := range strings.Split(word, "/") {

		if t.isWord {
			for k, _ := range t.next {
				if strings.Index(k, "{") != -1 {
					//除去{}字符，n表示替换n个，-1表示全部替换
					//让{id}脱壳，只存id字符串
					key := strings.Replace(k, "{", "", -1)
					key = strings.Replace(key, "}", "", -1)
					arg[key] = v
				}
				v = k
			}
		}
		if t.next[v] == nil {
			log.Println(word + "在路径未找到，路径树匹配不上")
			return
		}
		t = t.next[v]
	}
	if len(t.next) == 0 {
		isHave = t.isWord
		return
	}
	return
}

func (ctx *MyContext) Handler(url string, f func(context *MyContext)) {
	//路径注册
	//1、前缀树的插入，2、完成url到处理方法的映射
	UrlTree.Insert(url)
	ctx.routers[url] = f
}

func (ctx *MyContext) GetPathVariable(pathType, key string) string {
	path := ctx.Request.URL.Path
	if pathType == "" || pathType == "default" {
		return ctx.pathArgs[path][key]
	} else if strings.HasSuffix(path, "html") {
		path = strings.TrimSuffix(path, ".html")
		return ctx.pathArgs[path][key]
	}
	return ""
}

func (ctx *MyContext) GetForm(key string) (string, error) {
	if err := ctx.Request.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		return "", err
	}
	return ctx.Request.Form.Get(key), nil
}

func (ctx *MyContext) GetJson(key string) interface{} {
	var params map[string]interface{}
	body, _ := io.ReadAll(ctx.Request.Body)
	_ = json.Unmarshal(body, &params)
	return params[key]
}

// 实现了ServeHTTP接口
// listener监听时，net包在请求过来是会自动调用
// 作用：处理url跳转到对应方法
// 路径分三类：1、带.html 2、带/{} 2、前两个都不带
func (ctx *MyContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//信息写入上下文实体
	ctx.W = w
	ctx.Request = r
	path := r.URL.Path

	//查询path是否存在对应的处理函数
	//注意：静态路由地址在注册时一定能匹配上方法
	//    动态路由于{id}不同导致url不同，因此对应方法可能为空
	//    因此f为空时是处理动态路由，不为空时处理静态路由和已经存在的动态路由地址
	// 首先尝试匹配静态路由

	if handler, ok := ctx.routers[path]; ok {
		// 找到了静态路由，执行相应的处理函数
		handler(ctx)
		return
	}

	// 尝试匹配动态路由，首先分割路径
	parts := strings.Split(path, "/")

	// 如果路径以.html结尾，执行相应的处理逻辑
	if strings.Contains(path, "image") {
		// 提取图片路径
		for i, part := range parts {
			if part == "image" {
				// 从 "image" 后面的部分获取图片路径
				if i+1 < len(parts) {
					imagePath := strings.Join(parts[i+1:], "/")
					p := strings.Join(parts[:i+1], "/") + "/"
					if handler, ok := ctx.routers[p]; ok {

						// 找到了静态路由，执行相应的处理函数
						ctx.pathArgs[path][imagePath] = imagePath
						handler(ctx)

						return
					}
					return
				}
			}
		}
	}

	// 如果路径以.html结尾，则除去，交给动态路由匹配
	if strings.HasSuffix(path, ".html") {
		path = strings.Replace(path, ".html", "", -1)
	}
	// 尝试匹配动态路由
	for routersKey, handler := range ctx.routers {

		// 构建正则表达式模式，使用 .* 匹配变化的部分
		// 把routers中地址取出，把{id}改成{.*}用正则表达式匹配就行了！
		pattern := "^" + strings.Replace(routersKey, "{id}", ".*", -1) + "$"
		matched, err := regexp.MatchString(pattern, path)
		if err != nil {
			// 正则表达式匹配出错，可以进行错误处理
			http.Error(ctx.W, "正则表达式匹配出错", http.StatusInternalServerError)
			return
		}

		if matched {
			// 提取动态参数的值

			pathParams := strings.Split(path, "/")
			//re := regexp.MustCompile(pattern)
			//matches := re.FindStringSubmatch(path)
			if len(pathParams) > 1 {
				// 最后一个是开始是动态参数的值
				// 可以将动态参数的值提取并放入 ctx.pathArgs
				ctx.pathArgs[path] = map[string]string{"id": pathParams[len(pathParams)-1]}
			}

			// 匹配成功，执行相应的处理函数
			handler(ctx)
			return
		}
	}
}
