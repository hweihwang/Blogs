package blogstorage

import (
	"context"
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
)

func (s *sqlStore) Update(ctx context.Context, id uint, request *blogmodel.BlogUpdateRequest) error {
	return s.db.Where("id = ?", id).Updates(request).Error
}
