package component

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) AppContext {
	return &appCtx{
		db: db,
	}
}

func (appCtx *appCtx) GetMainDBConnection() *gorm.DB {
	return appCtx.db
}
