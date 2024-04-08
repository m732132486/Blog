package models

import (
	"time"

	"gorm.io/gorm"
)

// Params 注册参数
type Params struct {
	gorm.Model
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamsLogin 登录参数
type ParamsLogin struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamArticle 社区查询文章结构体
type ParamArticle struct {
	Username       string    `gorm:"column:username"`
	UserID         int64     `gorm:"user_id"`
	Title          string    `gorm:"column:article_title"`
	ParentID       int64     `gorm:"column:parent_id"`
	TitleId        int64     `gorm:"column:title_id"`
	ArticleContent string    `gorm:"column:article_content"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	CategoryName   string    `gorm:"column:category_name"`
	LikeCount      int64     `gorm:"column:like_count"`
}

type Query struct {
	Keyword string `json:"title"`
}

type CommunitySearchArticles struct {
	CategoryName string `json:"category_name"`
}

type UsernameSArticle struct {
	Page     int    `json:"page"`
	Username string `json:"username"`
}
