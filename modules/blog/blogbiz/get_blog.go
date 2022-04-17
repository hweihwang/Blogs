package blogbiz

import (
	"context"
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
)

type GetBlogStore interface {
	GetById(ctx context.Context, id uint) (*blogmodel.Blog, error)
}

type getBlogBiz struct {
	store GetBlogStore
}

func NewGetBlogBiz(store GetBlogStore) *getBlogBiz {
	return &getBlogBiz{
		store: store,
	}
}

func (b *getBlogBiz) GetBlog(ctx context.Context, id uint) (*blogmodel.Blog, error) {
	return b.store.GetById(ctx, id)
}
