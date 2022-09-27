package logic

import (
	"fmt"
	"strconv"

	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"cointiger.com/verification/common/convert"
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

func (l *UpdateverifyLogic) Updateverify(req *types.UpdateVerifyRequest, token string) (resp *types.UpdateVerifyResponse,
	err error) {
	resp = &types.UpdateVerifyResponse{
		IsExist: false,
	}
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
		//err = xerr.NewErr(constant.REPEATVERIFY, errors.New("已有此验证内容，不可添加"))
		resp.IsExist = true
		return
	}

	userName, strUserId := GetUserInfo(l.svcCtx, token)
	userId, _ := strconv.ParseInt(strUserId, 10, 64)

	data.VerifyInfo = req.VerifyInfo
	data.VerifyType = req.VerifyType
	data.SocialName = convert.StrToNullString(req.SocialName)
	data.JobTiele = convert.StrToNullString(req.JobTitle)
	data.Creator = convert.StrToNullString(userName)
	data.IsPay = convert.StrToNullString(req.IsPay)
	data.CreatorId = userId
	err = l.svcCtx.OfficialVerify.Update(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
	}
	return
}
