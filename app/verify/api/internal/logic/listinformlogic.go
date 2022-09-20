package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListinformLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListinformLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListinformLogic {
	return &ListinformLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListinformLogic) Listinform(req *types.ListInformRequest) (resp *types.ListInformResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
