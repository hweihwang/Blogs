package blogbiz

import (
	"context"
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
)

type UpdateBlogStore interface {
	GetById(ctx context.Context, id uint) (*blogmodel.Blog, error)

	Update(
		ctx context.Context,
		id uint,
		request *blogmodel.BlogUpdateRequest) error
}

type updateBlogBiz struct {
	store UpdateBlogStore
}

func NewUpdateBlogBiz(store UpdateBlogStore) *updateBlogBiz {
	return &updateBlogBiz{
		store: store,
	}
}

func (b *updateBlogBiz) UpdateBlog(ctx context.Context, id uint, request *blogmodel.BlogUpdateRequest) error {
	_, err := b.store.GetById(ctx, id)

	if err != nil {
		return err
	}

	return b.store.Update(ctx, id, request)
}
