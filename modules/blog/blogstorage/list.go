package blogstorage

import (
	"context"
	"github.com/hweihwang/go-blogs/common"
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
)

func (s *sqlStore) List(
	ctx context.Context,
	conditions map[string]interface{},
	filter *blogmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]*blogmodel.Blog, error) {

	var result []*blogmodel.Blog

	db := s.db

	db = db.Table((&blogmodel.Blog{}).TableName())

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	db = db.Where(conditions)

	if filter != nil {
		if filter.Title != nil {
			db = db.Where("title LIKE ?", filter.Title.GetLikeString())
		}
		if filter.Description != nil {
			db = db.Where("content LiKE ?", filter.Description.GetLikeString())
		}
		if filter.CreateById != nil {
			db = db.Where("author = ?", filter.CreateById)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset(paging.Limit * (paging.Page - 1)).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *sqlStore) ListSimple(
	ctx context.Context,
	filter *blogmodel.Filter,
	paging *common.Paging,
) ([]*blogmodel.Blog, error) {

	var results []*blogmodel.Blog

	db := s.db

	db = db.Table((&blogmodel.Blog{}).TableName())

	if filter != nil {
		if filter.Title != nil {
			db = db.Where("title LIKE ?", filter.Title.GetLikeString())
		}
		if filter.Description != nil {
			db = db.Where("content LiKE ?", filter.Description.GetLikeString())
		}
		if filter.CreateById != nil {
			db = db.Where("author = ?", filter.CreateById)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset(paging.Limit * (paging.Page - 1)).
		Limit(paging.Limit).
		Order("id desc").
		Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
