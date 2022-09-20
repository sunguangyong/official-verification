package handler

import (
	"net/http"

	"cointiger.com/verification/app/verify/api/internal/logic"
	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func exportinformHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListInformRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewExportinformLogic(r.Context(), svcCtx)
		resp, err := l.Exportinform(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
