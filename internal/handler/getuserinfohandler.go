package handler

import (
	"StreamAgent/response"
	"net/http"

	"StreamAgent/internal/logic"
	"StreamAgent/internal/svc"
	"StreamAgent/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func getUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserInfoLogic(r.Context(), ctx)
		resp, err := l.GetUserInfo(req)
		if err != nil {
			response.Response(w, resp, err)
		} else {
			response.Response(w, resp, nil)
		}
	}
}
