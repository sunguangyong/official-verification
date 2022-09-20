package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddverifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddverifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddverifyLogic {
	return &AddverifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddverifyLogic) Addverify(req *types.AddVerifyRequest) (resp *types.AddVerifyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
