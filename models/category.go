package models

type Category struct {
	Cid      int    `json:"cid" orm:"cid"`            // 分类ID
	Name     string `json:"name" orm:"name"`          // 分类名称
	CreateAt string `json:"createAt" orm:"create_at"` // 创建时间
	UpdateAt string `json:"updateAt" orm:"update_at"` // 更新时间
}

type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}
