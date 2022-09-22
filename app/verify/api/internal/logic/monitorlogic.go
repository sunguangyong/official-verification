package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MonitorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMonitorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MonitorLogic {
	return &MonitorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MonitorLogic) Monitor(req *types.MonitorRequest) (resp *types.MonitorResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
