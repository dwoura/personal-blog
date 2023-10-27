package service

import (
	"html/template"
	"personal-blog/config"
	"personal-blog/dao"
	"personal-blog/models"
	"regexp"
	"strings"
	"time"
)

func GetAllIndexInfo(slug string, page, pageSize int) (*models.HomeResponse, error) {
	//分类查询
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	//有slug的返回按slug查询
	if slug == "" {
		posts, err = dao.GetPostPage(page, pageSize)
	} else {
		posts, err = dao.GetPostBySlug(slug, page, pageSize)
	}
	//文章查询

	//把两个表中信息结合起来成一个实体
	var postMores []models.PostMore
	for _, post := range posts {
		categorysName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		//content内容可能过长，显示最大字数
		//rune为unicode字符
		content := []rune(post.Content)
		temp := string(content)
		//html标签匹配<p></p>中的内容
		re := regexp.MustCompile("<p>.*?</p>")
		match := re.FindString(temp)
		if match != "" {
			//除去换行符
			match = strings.ReplaceAll(match, "<p>", "")
			match = strings.ReplaceAll(match, "</p>", "")
			match = strings.ReplaceAll(match, "<br>", "")
			// 检查标签之前的字符数是否已经超过100个字符
			if len(match) > 100 {
				// 超过100个字符，保留前100个字符并保留标签
				match = match[:99]
				if strings.LastIndex(match, "<") != -1 {
					match = match[:strings.LastIndex(match, "<")]
				}

			}
			match = "<p>" + match + "...</p>"
			content = []rune(match)
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
	return hr, nil
}
