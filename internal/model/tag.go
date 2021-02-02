package model

// Tag model
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// TableName returns the name of tag table
func (t Tag) TableName() string {
	return "blog_tag"
}
