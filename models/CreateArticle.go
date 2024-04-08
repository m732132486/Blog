package models

import "time"

// Article 创建文章
type Article struct {
	ID             uint      `gorm:"primaryKey"`
	Title          string    `json:"title" binding:"required"` //标题
	TitleID        int64     `json:"title_id"`
	UserID         int64     `json:"user_id"`                                    //用户id
	ParentID       int64     `json:"parent_id" binding:"required,oneof=1 2 3 4"` //社区
	ArticleContent string    `json:"article_content" binding:"required"`         //文章内容
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

// ParamsArticle 查找文章
type ParamsArticle struct {
	Title          string `json:"title" binding:"required"` //文章名称
	Username       string `json:"username" binding:""`      //用户名字
	UserID         int64  `json:"user_id"`                  //用户id
	ParentID       int64  `json:"parent_id"`
	CategoryName   string `gorm:"column:category_name"` //所属社区
	TitleID        int64  `json:"title_id"`             //标题ID
	ArticleContent string `json:"article_content"`      //文章内容
	LikeCount      int64  `json:"like_count"`
}

// UserArticle 根据用户名查找文章
type UserArticle struct {
	Title          string `json:"title"`                       //文章名称
	Username       string `json:"username" binding:"required"` //用户名字
	UserID         int64  `json:"user_id"`                     //用户id
	ParentID       int64  `json:"parent_id"`                   //所属社区
	TitleID        int64  `json:"title_id"`                    //标题ID
	ArticleContent string `json:"article_content"`             //文章内容
	CategoryName   string `gorm:"column:category_name"`
	LikeCount      int64  `json:"like_count"`
}
type UserSArticle struct {
	Title          string `json:"title" binding:"required"` //文章名称
	ParentID       int64  `json:"parent_id"`                //所属社区
	ArticleContent string `json:"article_content"`          //文章内容
	CategoryName   string `gorm:"column:category_name"`
}

func (table *Article) TableName() string {
	return "articles"
}

type UserCollect struct {
	UserID    int64     `json:"user_id" binding:"required"`
	TitleID   int64     `json:"title_id"binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (table *UserCollect) TableName() string {
	return "user_collects"
}
