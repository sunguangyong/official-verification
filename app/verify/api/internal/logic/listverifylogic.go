package logic

import (
	"context"
	"fmt"
	"strings"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListverifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListverifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListverifyLogic {
	return &ListverifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListverifyLogic) Listverify(req *types.ListVerifyRequest) (resp *types.ListVerifyResponse, err error) {
	var querySql string
	queryList := make([]string, 0)
	verifyInfo := req.VerifyInfo
	verifyType := req.VerifyType
	socialName := req.SocialName

	resp.List = make([]types.ListVerify, 0)

	if verifyInfo != "" {
		queryList = append(queryList,fmt.Sprintf("verify_info like '%%%s%%'", verifyInfo))
	}

	if verifyType != "" {
		queryList = append(queryList,fmt.Sprintf("verify_type = '%s'", verifyType))
	}

	if socialName != "" {
		queryList = append(queryList,fmt.Sprintf("social_name = '%s'", socialName))
	}

	if len(queryList) > 0 {
		querySql = "where " + strings.Join(queryList," and ")
	}

	limitSql := fmt.Sprintf("limit %d, %d", (req.PageIndex-1)*req.PageSize, req.PageSize)
	orderSql := "order by id"

	count , err := l.svcCtx.OfficialVerify.FindNewsCount(l.ctx,querySql)

	if err != nil {
		l.Logger.Error(err)
		return
	}

	if count == 0 {
		return
	}

	resp.Count = int(count)
	dataList, err := l.svcCtx.OfficialVerify.CommonFind(l.ctx,querySql,limitSql,orderSql)

	if err != nil {
		l.Logger.Error(err)
		return
	}

	for _, data := range dataList {
		resp.List = append(resp.List,types.ListVerify{
			Id : int(data.Id),
			IsPay: data.IsPay.String,
			JobTiele: data.JobTiele.String,
			Creator: data.Creator.String,
			VerifyInfo:data.VerifyInfo,
			VerifyType: data.VerifyType,
			SocialName: data.SocialName.String,
			CreateTime: data.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return
}
