package handler

import (
	"net/http"

	"cointiger.com/verification/common/result"

	"cointiger.com/verification/app/verify/api/internal/logic"
	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func addverifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddVerifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		token := r.Header.Get("Authorization")
		l := logic.NewAddverifyLogic(r.Context(), svcCtx)
		resp, err := l.Addverify(&req, token)
		result.HttpResult(r, w, req, resp, err)
	}
}
