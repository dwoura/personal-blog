package dao

import (
	"log"
	"personal-blog/models"
)

func GetUser(userName, passwd string) (*models.User, error) {
	//从数据库中获取user信息并封装为User对象
	u := &models.User{} //对这个对象进行反射操作
	err := DB.QueryOne(u, "select * from blog_user where user_name=? and passwd=? limit 1", userName, passwd)
	return u, err
	//row := DB.QueryRow("select * from blog_user where user_name=? and passwd=? limit 1",
	//	userName,
	//	passwd)
	//if row.Err() != nil {
	//	log.Println(row.Err())
	//	return nil
	//}
	//var user = &models.User{}
	////把行数据输入到对象中
	//err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	//if err != nil {
	//	log.Println(row.Err())
	//	return nil
	//}
	//return user
}

func GetUserNameById(userId int) (userName string) {
	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	_ = row.Scan(&userName)
	return
}
