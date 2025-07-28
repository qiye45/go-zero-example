package svc

import (
	"bookstore/rpc/check/internal/config"
	"bookstore/rpc/model"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type ServiceContext struct {
	Config config.Config
	Model  model.BookModel // 手动代码

}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	// 添加MySQL心跳检测，每5分钟Ping一次
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			var one int
			// 用 context.Background()，也可以自定义超时
			err := conn.QueryRowCtx(context.Background(), &one, "SELECT 1")
			if err != nil {
				// 可以根据需要记录日志
			}
		}
	}()
	return &ServiceContext{
		Config: c,
		Model:  model.NewBookModel(conn, c.Cache), // 手动代码
	}
}
