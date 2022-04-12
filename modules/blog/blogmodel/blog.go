package blogmodel

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

type DateTime time.Time

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	if dt == nil {
		return []byte("null"), nil
	}
	return []byte(`"` + time.Time(*dt).Format("2006-01-02 15:04:05") + `"`), nil
}

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*dt = DateTime(t)
	return nil
}

func (dt *DateTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return errors.New("invalid type for DateTime")
	}
	*dt = DateTime(t)
	return nil
}

func (dt DateTime) Value() (driver.Value, error) {
	return time.Time(dt), nil
}

type Blog struct {
	BlogCreateRequest
}

func (*Blog) TableName() string {
	return "blogs"
}

type BlogCreateRequest struct {
	Id          uint       `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title       string     `json:"title" gorm:"column:title;index:idx_title;not null"`
	Description *string    `json:"description" gorm:"column:description;index:idx_description"`
	Content     *string    `json:"content" gorm:"column:content;index:idx_content"`
	CreateById  *uint      `json:"create_by_id" gorm:"column:create_by_id;index:idx_create_by_id"`
	CreatedAt   DateTime  `json:"created_at" gorm:"column:created_at;index:idx_created_at"`
	UpdatedAt   *DateTime `json:"updated_at" gorm:"column:updated_at;index:idx_updated_at"`
}

func (*BlogCreateRequest) TableName() string {
	return (&Blog{}).TableName()
}

func (r *BlogCreateRequest) Validate() error {
	r.Title = strings.TrimSpace(r.Title)

	if r.Title == "" {
		return errors.New("title is required")
	}

	return nil
}
