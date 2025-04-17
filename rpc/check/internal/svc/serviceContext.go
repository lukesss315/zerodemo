package svc

import (
	"demo.exmale.com/model/bookstore"
	"demo.exmale.com/rpc/check/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  bookstore.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  bookstore.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
