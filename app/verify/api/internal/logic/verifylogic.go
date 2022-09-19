package logic

import (
	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"cointiger.com/verification/app/verify/model"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLogic) Verify(req *types.AddVerifyRequest) (resp *types.AddVerifyResponse, err error) {
	resp = &types.AddVerifyResponse{
		Repetition: make([]string, 0),
	}

	verifyType := req.VerifyType
	verifyInfos := req.VerifyInfo
	jobTiele := req.JobTiele
	isPay := req.IsPay
	creator := req.Creator

	for _, info := range verifyInfos {
		querySql := fmt.Sprintf("where verify_info = '%s'", info)
		arrys, _ := l.svcCtx.OfficialVerify.CommonFind(l.ctx, querySql, "", "limit 1")
		if len(arrys) > 0 {
			resp.Repetition = append(resp.Repetition, info)
		}
	}

	if len(resp.Repetition) > 0 {
		return
	}

	for _, info := range verifyInfos {
		officialVerify := &model.OfficialVerify{
			VerifyType: verifyType,
			VerifyInfo: info,
			JobTiele:   jobTiele,
			IsPay:      isPay,
			Creator:    creator,
		}
		l.svcCtx.OfficialVerify.Insert(l.ctx, officialVerify)
	}
	return
}
