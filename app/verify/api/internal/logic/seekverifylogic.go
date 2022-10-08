package logic

import (
	"cointiger.com/verification/app/verify/model"
	"cointiger.com/verification/common/xerr"
	"context"
	"fmt"
	"strings"

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

func (l *SeekverifyLogic) Seekverify(req *types.SeekVerifyRequest, language string) (resp *types.SeekVerifyResponse,
	err error) {
	resp = &types.SeekVerifyResponse{
		VerifyList: make([]types.ListVerify, 0),
	}

	if len(req.VerifyInfo) > 1024 {
		return nil, xerr.NewParamsErr(language)
	}

	if len(req.VerifyType) > 50 {
		return nil, xerr.NewParamsErr(language)
	}

	var listVerify []*model.OfficialVerify
	switch req.VerifyType {
	case "Website":
		listVerify, err = l.website(req)

	case "Telegram Username":
		listVerify, err = l.telegramUsername(req)

	case "Social media":
		listVerify, err = l.common(req)

	case "Discord ID":
		listVerify, err = l.common(req)
	}

	if err != nil {
		return
	}

	for _, data := range listVerify {
		var Verify types.ListVerify
		Verify.VerifyType = data.VerifyType
		Verify.VerifyInfo = data.VerifyInfo
		Verify.SocialName = data.SocialName.String
		Verify.JobTiele = data.JobTiele.String
		Verify.IsPay = data.IsPay.String
		Verify.CreateTime = data.CreateTime.Format("2006-01-02 15:04:05")
		Verify.Creator = data.Creator.String
		Verify.Id = int(data.Id)
		resp.VerifyList = append(resp.VerifyList, Verify)
	}

	return
}

func (l *SeekverifyLogic) website(req *types.SeekVerifyRequest) (listVerify []*model.OfficialVerify, err error) {
	listVerify = make([]*model.OfficialVerify, 0)
	querySql := fmt.Sprintf("where verify_type = '%s'", req.VerifyType)
	verifyList, err := l.svcCtx.OfficialVerify.CommonFind(l.ctx, querySql, "", "")

	if err != nil {
		l.Logger.Error(err)
		return
	}

	for _, verify := range verifyList {
		ok := strings.Contains(req.VerifyInfo, verify.VerifyInfo)
		if ok {
			listVerify = append(listVerify, verify)
		}
	}

	return
}

func (l *SeekverifyLogic) common(req *types.SeekVerifyRequest) (listVerify []*model.OfficialVerify, err error) {
	listVerify = make([]*model.OfficialVerify, 0)
	querySql := fmt.Sprintf("where verify_type = '%s' and verify_info ='%s' ", req.VerifyType, req.VerifyInfo)
	verifyList, err := l.svcCtx.OfficialVerify.CommonFind(l.ctx, querySql, "", "")

	if err != nil {
		l.Logger.Error(err)
		return
	}

	for _, verify := range verifyList {
		listVerify = append(listVerify, verify)
	}

	return
}

func (l *SeekverifyLogic) telegramUsername(req *types.SeekVerifyRequest) (listVerify []*model.OfficialVerify,
	err error) {
	listVerify = make([]*model.OfficialVerify, 0)
	var verifyInfo string

	ok := strings.HasPrefix(req.VerifyInfo, "@")
	if ok {
		verifyInfo = req.VerifyInfo[1:]
	} else {
		verifyInfo = req.VerifyInfo
	}

	querySql := fmt.Sprintf("where verify_type = '%s' and verify_info ='%s' ", req.VerifyType, verifyInfo)
	verifyList, err := l.svcCtx.OfficialVerify.CommonFind(l.ctx, querySql, "", "")

	if err != nil {
		l.Logger.Error(err)
		return
	}

	for _, verify := range verifyList {
		listVerify = append(listVerify, verify)
	}

	return
}
