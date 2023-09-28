package handler

import (
	"net/http"

	"google.golang.org/grpc/metadata"

	"cointiger.com/verification/common/result"

	"cointiger.com/verification/app/verify/api/internal/logic"
	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func listinformHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListInformRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		// 将http token 传入 rpc metadata
		md := metadata.Pairs()
		for key, values := range r.Header {
			if key == "Authorization" {
				if len(values) > 0 {
					md.Append(key, values[0])
				}
			}
		}

		ctx := metadata.NewOutgoingContext(r.Context(), md)
		l := logic.NewListinformLogic(ctx, svcCtx)

		resp, err := l.Listinform(&req)
		result.HttpResult(r, w, req, resp, err)
	}
}
