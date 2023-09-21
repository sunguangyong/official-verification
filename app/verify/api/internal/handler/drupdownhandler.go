package handler

import (
	"net/http"

	"cointiger.com/verification/common/result"

	"cointiger.com/verification/app/verify/api/internal/logic"
	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func drupdownHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DropdownRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDrupdownLogic(r.Context(), svcCtx)
		resp, err := l.Drupdown(&req)
		result.HttpResult(r, w, req, resp, err)
	}
}
