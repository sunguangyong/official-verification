package logic

import (
	"cointiger.com/verification/app/verify/model"
	"context"
	"fmt"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"cointiger.com/verification/common/convert"

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
	resp = &types.AddVerifyResponse{
		Repetition: make([]types.VerifyData, 0),
	}

	verifyType := req.VerifyType
	dataList := req.Data

	for _, info := range dataList {
		querySql := fmt.Sprintf("where verify_info = '%s'", info.VerifyInfo)
		arrys, _ := l.svcCtx.OfficialVerify.CommonFind(l.ctx, querySql, "", "limit 1")
		if len(arrys) > 0 {
			resp.Repetition = append(resp.Repetition, info)
		}
	}

	if len(resp.Repetition) > 0 {
		return
	}

	for _, info := range dataList {
		officialVerify := &model.OfficialVerify{
			VerifyType: verifyType,
			VerifyInfo: info.VerifyInfo,
			JobTiele:   convert.StrToNullString(info.JobTitle),
			IsPay:      convert.StrToNullString(info.IsPay),
			Creator:    convert.StrToNullString(req.Creator),
			SocialName: convert.StrToNullString(req.SocialName),
		}
		l.svcCtx.OfficialVerify.Insert(l.ctx, officialVerify)
	}
	return
}
