package handler

import (
	"fmt"
	"net/http"
	"time"

	"cointiger.com/verification/common/excel"

	"cointiger.com/verification/app/verify/api/internal/logic"
	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func exportinformHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExportInformRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewExportinformLogic(r.Context(), svcCtx)
		resp, err := l.Exportinform(&req)

		data := make([]interface{}, 0)
		if len(resp.List) > 0 && err == nil {
			for _, v := range resp.List {
				data = append(data, v)
			}
		}
		disposition := fmt.Sprintf("attachment; filename=%s.xlsx", time.Now().Format("2006-01-02-15-04-05"))
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", disposition)
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		excel.ExportByStruct(w, []string{"举报内容", "类型", "名称", "时间"}, data, "b")
		return
	}
}
