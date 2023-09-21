package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"cointiger.com/verification/app/verify/grpc/verify"

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
	data, err := l.svcCtx.VerifyRpc.ListInform(l.ctx, &verify.ListInformRequest{
		VerifyInfo: req.VerifyInfo,
		PageIndex:  req.PageIndex,
		PageSize:   req.PageSize,
		VerifyType: req.VerifyType,
		SocialName: req.SocialName,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
	})

	resp = &types.ListInformResponse{
		List:  make([]types.ListInform, 0),
		Count: int(data.Count),
	}

	for _, v := range data.List {
		resp.List = append(resp.List, types.ListInform{
			VerifyInfo: v.VerifyInfo,
			VerifyType: v.VerifyType,
			SocialName: v.SocialName,
			CreateTime: v.CreateTime,
		})
	}

	return
}
