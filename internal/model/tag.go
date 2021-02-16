package model

import "github.com/baconYao/gin-blog/pkg/app"

// Tag model
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// TagSwagger is a struct for Swagger API doc
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// TableName returns the name of tag table
func (t Tag) TableName() string {
	return "blog_tag"
}
