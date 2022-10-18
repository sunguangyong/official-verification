package logic

import (
	"context"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DrupdownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDrupdownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrupdownLogic {
	return &DrupdownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DrupdownLogic) Drupdown(req *types.DropdownRequest) (resp *types.DropdownResponse, err error) {
	resp = &types.DropdownResponse{
		SocialDropdown: make([]types.Dropdown, 0),
		VerifyDropdown: make([]types.Dropdown, 0),
	}
	//verifyList, _ := l.svcCtx.VerifyEnum.CommonFind(l.ctx, "where is_delete = 0 ", "", "")
	socialList, _ := l.svcCtx.SocialEnum.CommonFind(l.ctx, "where is_delete = 0 ", "", "")

	//for _, data := range verifyList {
	//	resp.VerifyDropdown = append(resp.VerifyDropdown, types.Dropdown{
	//		Label: data.VerifyType,
	//		Value: data.VerifyType,
	//	})
	//}

	resp.VerifyDropdown = []types.Dropdown {
		types.Dropdown{
			Value: "Website",
			Label: "Website Verification",
		},
		types.Dropdown{
			Value: "Social media",
			Label: "Social media",
		},
		types.Dropdown{
			Value: "Telegram Username",
			Label: "Telegram Account",
		},
		types.Dropdown{
			Value: "Discord ID",
			Label: "Discord Account",
		},
	}

	for _, data := range socialList {
		resp.SocialDropdown = append(resp.SocialDropdown, types.Dropdown{
			Label: data.SocialName,
			Value: data.SocialName,
		})
	}

	return
}
