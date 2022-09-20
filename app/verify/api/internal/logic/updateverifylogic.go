package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateverifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateverifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateverifyLogic {
	return &UpdateverifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateverifyLogic) Updateverify(req *types.UpdateVerifyRequest) (resp *types.AddVerifyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
