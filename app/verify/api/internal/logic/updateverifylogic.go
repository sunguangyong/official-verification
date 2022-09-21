package logic

import (
	"fmt"

	"errors"

	"context"

	"cointiger.com/verification/common/constant"
	"cointiger.com/verification/common/convert"
	"cointiger.com/verification/common/xerr"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateverifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateverifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateverifyLogic {
	return &UpdateverifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateverifyLogic) Updateverify(req *types.UpdateVerifyRequest) (resp *types.UpdateVerifyResponse, err error) {
	id := req.Id
	data, err := l.svcCtx.OfficialVerify.FindOne(l.ctx, id)
	if data == nil || err != nil {
		l.Logger.Error(err)
		return
	}
	querySql := fmt.Sprintf("where verify_info = '%s' and id != %d", req.VerifyInfo, id)
	verifyList, err := l.svcCtx.OfficialVerify.CommonFind(l.ctx, querySql, "", "limit 1")

	if err != nil {
		l.Logger.Error(err)
		return
	}

	if len(verifyList) > 0 {
		err = xerr.NewErr(constant.REPEATVERIFY, errors.New("已有此验证内容，不可添加"))
		return
	}

	data.VerifyInfo = req.VerifyInfo
	data.VerifyType = req.VerifyType
	data.SocialName = convert.StrToNullString(req.SocialName)
	data.JobTiele = convert.StrToNullString(req.JobTitle)
	data.Creator = convert.StrToNullString(req.Creator)
	data.IsPay = convert.StrToNullString(req.IsPay)
	err = l.svcCtx.OfficialVerify.Update(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
	}
	return
}
