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
	id := req.Id
	err = l.svcCtx.OfficialVerify.Delete(l.ctx, id)
	if err != nil {
		l.Logger.Error(err)
	}
	return
}
