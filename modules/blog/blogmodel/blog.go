package blogmodel

import (
	"errors"
	"github.com/hweihwang/go-blogs/common"
	"strings"
)

type Blog struct {
	common.SQLModel `json:",inline"`
	Title           string  `json:"title" gorm:"column:title;index:idx_title;not null"`
	Description     *string `json:"description" gorm:"column:description;index:idx_description"`
	Content         *string `json:"content" gorm:"column:content;index:idx_content"`
	CreatedById     uint   `json:"created_by_id" gorm:"column:created_by_id;index:idx_create_by_id"`
}

func (*Blog) TableName() string {
	return "blogs"
}

type BlogCreateRequest struct {
	common.SQLModel `json:",inline"`
	Title           string  `json:"title" gorm:"column:title;index:idx_title;not null"`
	Description     *string `json:"description" gorm:"column:description;index:idx_description"`
	Content         *string `json:"content" gorm:"column:content;index:idx_content"`
	CreatedById     *uint   `json:"created_by_id" gorm:"column:created_by_id;index:idx_create_by_id"`
}

type BlogUpdateRequest struct {
	common.SQLModel `json:",inline"`
	Title           *string  `json:"title" gorm:"column:title;index:idx_title;not null"`
	Description     *string `json:"description" gorm:"column:description;index:idx_description"`
	Content         *string `json:"content" gorm:"column:content;index:idx_content"`
	CreatedById     *uint   `json:"created_by_id" gorm:"column:created_by_id;index:idx_create_by_id"`
}

func (*BlogCreateRequest) TableName() string {
	return (&Blog{}).TableName()
}

func (*BlogUpdateRequest) TableName() string {
	return (&Blog{}).TableName()
}

func (r *BlogCreateRequest) Validate() error {
	r.Title = strings.TrimSpace(r.Title)

	if r.Title == "" {
		return errors.New("title is required")
	}

	return nil
}
