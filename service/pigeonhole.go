package service

import (
	"personal-blog/config"
	"personal-blog/dao"
	"personal-blog/models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	//查询所有的文章 进行月份的整理
	//查询所有的分类
	posts, _ := dao.GetAllPost()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		//把所有文章按各自月份啊append到map[年月]的post数组中
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	categorys, _ := dao.GetAllCategory()
	return models.PigeonholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		pigeonholeMap,
	}
}
