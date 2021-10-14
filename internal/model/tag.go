package model

import "go-blog-s/pkg/app"

type Tag struct {
	//Model* Model
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (tag Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
