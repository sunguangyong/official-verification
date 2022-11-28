package svc

import (
	"cointiger.com/verification/app/verify/api/internal/config"
	"cointiger.com/verification/app/verify/model"
	"cointiger.com/verification/common/random"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	VeriflyRedis   *redis.Redis
	OfficialVerify model.OfficialVerifyModel
	VerifyEnum     model.VerifyEnumModel
	SocialEnum     model.SocialEnumModel
	ReportRecord   model.ReportRecordModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		VeriflyRedis:   redis.New(random.RandomStringArry(c.VerifyRdb.Addr), redis.WithPass(c.VerifyRdb.Passwd)),
		OfficialVerify: model.NewOfficialVerifyModel(sqlx.NewMysql(c.VerifyMysql.DataSource)),
		VerifyEnum:     model.NewVerifyEnumModel(sqlx.NewMysql(c.VerifyMysql.DataSource)),
		SocialEnum:     model.NewSocialEnumModel(sqlx.NewMysql(c.VerifyMysql.DataSource)),
		ReportRecord:   model.NewReportRecordModel(sqlx.NewMysql(c.VerifyMysql.DataSource)),
	}
}
