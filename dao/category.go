package dao

import (
	"log"
	"personal-blog/models"
)

func GetAllCategory() ([]models.Category, error) {
	//查询多行用DB.Query,rows变量用.Next()遍历
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 取值出错:", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}

func GetCategoryNameById(cId int) (categoryName string) {
	row := DB.QueryRow("select name from blog_category where cid=?", cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	_ = row.Scan(&categoryName)
	return
}
