package model

import "github.com/baconYao/gin-blog/pkg/app"

// Article model
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

// ArticleSwagger is a struct for Swagger API doc
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

// TableName returns the name of article table
func (a Article) TableName() string {
	return "blog_article"
}
