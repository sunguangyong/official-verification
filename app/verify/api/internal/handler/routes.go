// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"cointiger.com/verification/app/verify/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/offical/verify/dropdown",
				Handler: drupdownHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/offical/verify/add",
				Handler: addverifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/offical/verify/listverify",
				Handler: listverifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/offical/verify/detail",
				Handler: detailverifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/offical/verify/update",
				Handler: updateverifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/offical/verify/delete",
				Handler: deleteverifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/offical/verify/listinform",
				Handler: listinformHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/offical/verify/exportinform",
				Handler: exportinformHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/offical/verify/informadd",
				Handler: addinformHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/offical/verify/seekverify",
				Handler: seekverifyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/offical/verify/monitor",
				Handler: monitorHandler(serverCtx),
			},
		},
	)
}
