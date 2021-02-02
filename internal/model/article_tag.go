package model

// ArticleTag model
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

// TableName returns the name of article_tag table
func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
