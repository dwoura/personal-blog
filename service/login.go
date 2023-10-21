package service

import (
	"errors"
	"personal-blog/dao"
	"personal-blog/models"
	"personal-blog/utils"
)

//登录处理

func Login(userName, passwd string) (*models.LoginRes, error) {
	//前端和后端都了加密一次
	passwd = utils.Md5Crypt(passwd, "myblog")
	user, err := dao.GetUser(userName, passwd)
	if user == nil {
		//登录失败
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	//生成token 用到jwt技术进行生成 令牌，用来在一段时间内不用通过登录获取数据
	//jwt有三部分A.B.C 自行了解
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
