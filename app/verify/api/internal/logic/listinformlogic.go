package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"cointiger.com/verification/app/verify/grpc/verify"
	"cointiger.com/verification/common/convert"
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
	var request = &verify.ListInformRequest{}

	resp = &types.ListInformResponse{
		List:  make([]types.ListInform, 0),
		Count: 0,
	}

	convert.CopyProperties(request, req)

	data, err := l.svcCtx.VerifyRpc.ListInform(l.ctx, request)

	if err != nil {
		return
	}

	convert.CopyProperties(resp, data)

	return
}
