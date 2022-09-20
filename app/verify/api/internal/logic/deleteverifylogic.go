package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteverifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteverifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteverifyLogic {
	return &DeleteverifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteverifyLogic) Deleteverify(req *types.DeleteVerifyRequest) (resp *types.DeleteVerifyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
