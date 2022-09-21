package logic

import (
	"context"
	"fmt"
	"strings"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

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
	var querySql string
	queryList := make([]string, 0)
	verifyInfo := req.VerifyInfo
	verifyType := req.VerifyType
	socialName := req.SocialName
	createTime := req.StartTime
	endTime := req.EndTime

	resp = &types.ListInformResponse{
		List:  make([]types.ListInform, 0),
		Count: 0,
	}

	resp.List = make([]types.ListInform, 0)

	if verifyInfo != "" {
		queryList = append(queryList, fmt.Sprintf("verify_info like '%%%s%%'", verifyInfo))
	}

	if verifyType != "" {
		queryList = append(queryList, fmt.Sprintf("verify_type = '%s'", verifyType))
	}

	if socialName != "" {
		queryList = append(queryList, fmt.Sprintf("social_name = '%s'", socialName))
	}

	if createTime != "" && endTime != "" {
		queryList = append(queryList, fmt.Sprintf("create_time between '%s' and '%s'", createTime, endTime))
	}

	if len(queryList) > 0 {
		querySql = "where " + strings.Join(queryList, " and ")
	}

	limitSql := fmt.Sprintf("limit %d, %d", (req.PageIndex-1)*req.PageSize, req.PageSize)
	orderSql := "order by id"

	count, err := l.svcCtx.ReportRecord.FindNewsCount(l.ctx, querySql)

	if err != nil {
		l.Logger.Error(err)
		return
	}

	if count == 0 {
		return
	}

	resp.Count = int(count)
	dataList, err := l.svcCtx.ReportRecord.CommonFind(l.ctx, querySql, orderSql, limitSql)

	if err != nil {
		l.Logger.Error(err)
		return
	}

	for _, data := range dataList {
		resp.List = append(resp.List, types.ListInform{
			VerifyInfo: data.VerifyInfo,
			VerifyType: data.VerifyType,
			SocialName: data.SocialName.String,
			CreateTime: data.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return
}
