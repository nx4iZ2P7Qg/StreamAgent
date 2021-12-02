// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"StreamAgent/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/srs/dvr",
				Handler: DvrHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/login",
				Handler: LoginHandler(serverCtx),
			},
		},
	)
}
