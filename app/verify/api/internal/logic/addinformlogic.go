package logic

import (
	"cointiger.com/verification/app/verify/model"
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"

	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddinformLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddinformLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddinformLogic {
	return &AddinformLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddinformLogic) Addinform(req *types.SeekVerifyRequest) (resp *types.SeekVerifyResponse, err error) {
	var report model.ReportRecord
	report.VerifyInfo = req.VerifyInfo
	report.VerifyType = req.VerifyType
	_, err = l.svcCtx.ReportRecord.Insert(l.ctx, &report)
	if err != nil {
		l.Logger.Error(err)
	}
	return
}
