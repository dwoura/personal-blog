package service

import (
	"html/template"
	"personal-blog/config"
	"personal-blog/dao"
	"personal-blog/models"
	"time"
)

func GetPostsByCategoryId(cId, page, pageSize int) (*models.CategoryResponse, error) {
	//分类查询
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	//文章查询
	posts, err := dao.GetPostPageByCategoryId(cId, page, pageSize)
	//把两个表中信息结合起来成一个实体
	var postMores []models.PostMore
	for _, post := range posts {
		categorysName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		//content内容可能过长，显示最大字数
		//rune为unicode字符
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categorysName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			post.CreateAt.Format(time.DateTime),
			post.UpdateAt.Format(time.DateTime),
		}
		//把视图添加到页面集去
		postMores = append(postMores, postMore)
	}
	//分页
	//例：若条数共11 当页大小10要2页
	//算页的表达式 为(total-1)/pageSize + 1,例如(11-1)/10+1=2
	// /有向下取整的作用，加上1必定能把所有条数包含进页中
	total := dao.CountGetAllPost()
	//一页十条
	pagesCount := (total-1)/10 + 1 //页数
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page == pagesCount, //判断页号是否到页数（最后），最后一页要特殊处理
	}
	categoryName := dao.GetCategoryNameById(cId)
	categoryResponse := &models.CategoryResponse{
		hr,
		categoryName,
	}
	return categoryResponse, nil
}
