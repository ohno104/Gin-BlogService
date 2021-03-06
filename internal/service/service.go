package service

import (
	"context"

	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/global"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/internal/dao"
	otgorm "github.com/eddycjy/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	//svc.dao = dao.New(global.DBEngine)
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
