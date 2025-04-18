package logic

import (
	"context"
	"demo.exmale.com/model/bookstore"

	"demo.exmale.com/rpc/add/add"
	"demo.exmale.com/rpc/add/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.Model.Insert(l.ctx, &bookstore.Book{
		Book:  in.Book,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}
	return &add.AddResp{
		Ok: true,
	}, nil
}
