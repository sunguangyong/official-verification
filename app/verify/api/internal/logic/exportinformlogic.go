package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportinformLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportinformLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportinformLogic {
	return &ExportinformLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportinformLogic) Exportinform(req *types.ListInformRequest) (resp *types.ExportInformResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
