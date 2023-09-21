package logic

import (
	"context"
	"fmt"
	"strings"

	"cointiger.com/verification/app/verify/grpc/internal/svc"
	"cointiger.com/verification/app/verify/grpc/verify"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListInformLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListInformLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListInformLogic {
	return &ListInformLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListInformLogic) ListInform(in *verify.ListInformRequest) (*verify.ListInformResponse, error) {
	var querySql string
	queryList := make([]string, 0)
	verifyInfo := in.VerifyInfo
	verifyType := in.VerifyType
	socialName := in.SocialName
	createTime := in.StartTime
	endTime := in.EndTime

	resp := &verify.ListInformResponse{
		List:  make([]*verify.ListInform, 0),
		Count: 0,
	}

	resp.List = make([]*verify.ListInform, 0)

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

	limitSql := fmt.Sprintf("limit %d, %d", (in.PageIndex-1)*in.PageSize, in.PageSize)
	orderSql := "order by id desc"

	count, err := l.svcCtx.ReportRecord.FindNewsCount(l.ctx, querySql)

	if err != nil {
		l.Logger.Error(err)
		return resp, err
	}

	fmt.Println("count === ", count)

	if count == 0 {
		return resp, err
	}

	resp.Count = int32(count)
	dataList, err := l.svcCtx.ReportRecord.CommonFind(l.ctx, querySql, orderSql, limitSql)

	if err != nil {
		l.Logger.Error(err)
		return resp, err
	}

	for _, data := range dataList {
		resp.List = append(resp.List, &verify.ListInform{
			VerifyInfo: data.VerifyInfo,
			VerifyType: data.VerifyType,
			SocialName: data.SocialName.String,
			CreateTime: data.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return resp, nil
}
