package service

import (
	"context"
	"github.com/DATOULIN/dtservice/internal/dtservice/dao"
	"github.com/DATOULIN/dtservice/internal/pkg/setting"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(setting.DBEngine)
	return svc
}
