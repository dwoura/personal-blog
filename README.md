# personal-blog个人博客系统v1.1
## 更新内容
1、修复了文章的删除功能，以及上传的图片无法显示的问题。<br>
2、页脚备案号已添加。<br>
3、首页、分类显示文章简介不再受内容中html标签影响。<br>
4、重新组织了router，使用基于前缀树的动态路由匹配。<br>
## go无框架原生个人博客
基本架构：Golang 简洁架构。 类似javaweb的mvc三层,models、views和controller<br>
博客展示:www.dwoura.top
## 在mvc三层中对应的文件夹
+ models: 状态改变(一般是业务逻辑)<br>
service、dao
+ views: 绑定、展示m层数据，提供可交互ui<br>
views、template
+ controller: 接收用户请求、委托m层进行处理、处理数据返回给v层<br>
router
+ 剩下的文件夹:<br>
api: 封装界面请求接口<br>
common: 封装了一部分全局函数<br>
config: 配置文件及读取<br>
public: 全局资源<br>
server: 封装程序启动代码(减少了main中代码)<br>
utils: 封装工具类(加密、验证必备)<br>

## Golang 简洁架构（Clean Architecture）介绍
+ 外部依赖层（External Dependencies）
+ 实体层（Entity Layer）
+ 用例层（Use Case Layer
+ 接口适配器层（Interface Adapters Layer）
+ 主程序（Main）
