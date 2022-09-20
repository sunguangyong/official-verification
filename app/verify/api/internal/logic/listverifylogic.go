package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListverifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListverifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListverifyLogic {
	return &ListverifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListverifyLogic) Listverify(req *types.ListVerifyRequest) (resp *types.AddVerifyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
