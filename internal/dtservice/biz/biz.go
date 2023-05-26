package biz

import (
	"context"
	"github.com/DATOULIN/dtservice/internal/dtservice/dao"
	"github.com/DATOULIN/dtservice/internal/pkg/helper"
)

type Biz struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Biz {
	biz := Biz{ctx: ctx}
	biz.dao = dao.New(helper.DBEngine)
	return biz
}
