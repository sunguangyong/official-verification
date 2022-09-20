package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DroupdownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDroupdownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DroupdownLogic {
	return &DroupdownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DroupdownLogic) Droupdown(req *types.DropdownRequest) (resp *types.DropdownResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
