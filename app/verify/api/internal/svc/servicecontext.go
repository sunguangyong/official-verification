package svc

import (
	"cointiger.com/verification/app/verify/api/internal/config"
	"cointiger.com/verification/app/verify/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	OfficialVerify model.OfficialVerifyModel
	VerifyEnum     model.VerifyEnumModel
	SocialEnum     model.SocialEnumModel
	ReportRecord   model.ReportRecordModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		OfficialVerify: model.NewOfficialVerifyModel(sqlx.NewMysql(c.VerifyMysql.DataSource)),
		VerifyEnum:     model.NewVerifyEnumModel(sqlx.NewMysql(c.VerifyMysql.DataSource)),
		SocialEnum:     model.NewSocialEnumModel(sqlx.NewMysql(c.VerifyMysql.DataSource)),
		ReportRecord:   model.NewReportRecordModel(sqlx.NewMysql(c.VerifyMysql.DataSource)),
	}
}
