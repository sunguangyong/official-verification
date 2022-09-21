package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailverifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailverifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailverifyLogic {
	return &DetailverifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailverifyLogic) Detailverify(req *types.DetailVerifyRequest) (resp *types.DetailVerifyResponse, err error) {
	resp = &types.DetailVerifyResponse{
		Verify: types.ListVerify{},
	}
	data, err := l.svcCtx.OfficialVerify.FindOne(l.ctx, req.Id)
	if err != nil {
		l.Logger.Error(err)
	}

	resp.Verify.VerifyType = data.VerifyType
	resp.Verify.VerifyInfo = data.VerifyInfo
	resp.Verify.SocialName = data.SocialName.String
	resp.Verify.JobTiele = data.JobTiele.String
	resp.Verify.IsPay = data.IsPay.String
	resp.Verify.CreateTime = data.CreateTime.Format("2006-01-02 15:04:05")
	resp.Verify.Creator = data.Creator.String
	resp.Verify.Id = int(data.Id)

	return
}
