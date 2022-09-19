package svc

import (
	"cointiger.com/verification/app/verify/api/internal/config"
	"cointiger.com/verification/app/verify/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	OfficialVerify model.OfficialVerifyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		OfficialVerify: model.NewOfficialVerifyModel(sqlx.NewMysql(c.VerifyMysql.DataSource)),
	}
}
