package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeekverifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeekverifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeekverifyLogic {
	return &SeekverifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeekverifyLogic) Seekverify(req *types.SeekVerifyRequest) (resp *types.SeekVerifyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
