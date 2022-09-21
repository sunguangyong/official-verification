package handler

import (
	"cointiger.com/verification/common/result"
	"net/http"

	"cointiger.com/verification/app/verify/api/internal/logic"
	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func listverifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListVerifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListverifyLogic(r.Context(), svcCtx)
		resp, err := l.Listverify(&req)
		result.HttpResult(r, w, resp, err)
	}
}
