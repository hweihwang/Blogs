package blogstorage

import (
	"context"
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
)

func (s *sqlStore) GetById(
	ctx context.Context,
	id int64,
) (*blogmodel.Blog, error) {
	var b blogmodel.Blog

	db := s.db

	if err := db.First(&b, id).Error; err != nil {
		return nil, err
	}

	return &b, nil
}
