// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"Gopan/app/upload/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JWT},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: fileUploadHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/file"),
	)
}