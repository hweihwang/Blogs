package blogstorage

import (
	"context"
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
)

func (s *sqlStore) Create(ctx context.Context, request *blogmodel.BlogCreateRequest) error {
	return s.db.Create(request).Error
}
