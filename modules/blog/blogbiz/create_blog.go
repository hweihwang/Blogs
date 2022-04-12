package blogbiz

import (
	"context"
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
)

type CreateBlogStore interface {
	Create(ctx context.Context, request *blogmodel.BlogCreateRequest) error
}

type createBlogBiz struct {
	store CreateBlogStore
}

func NewCreateBlogBiz(store CreateBlogStore) *createBlogBiz {
	return &createBlogBiz{
		store: store,
	}
}

func (b *createBlogBiz) CreateBlog(ctx context.Context, request *blogmodel.BlogCreateRequest) error {
	if err := request.Validate(); err != nil {
		return err
	}

	return b.store.Create(ctx, request)
}
