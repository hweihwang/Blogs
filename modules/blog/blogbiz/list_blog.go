package blogbiz

import (
	"context"
	"github.com/hweihwang/go-blogs/common"
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
)

type ListBlogStore interface {
	List(
		ctx context.Context,
		conditions map[string]interface{},
		filter *blogmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]*blogmodel.Blog, error)

	ListSimple(
		ctx context.Context,
		filter *blogmodel.Filter,
		paging *common.Paging,
	) ([]*blogmodel.Blog, error)
}

type listBlogBiz struct {
	store ListBlogStore
}

func NewListBlogBiz(store ListBlogStore) *listBlogBiz {
	return &listBlogBiz{
		store: store,
	}
}

func (b *listBlogBiz) ListBlog(
	ctx context.Context,
	filter *blogmodel.Filter,
	paging *common.Paging,
) ([]*blogmodel.Blog, error) {
	return b.store.ListSimple(ctx, filter, paging)
}
