package handler

import (
	"StreamAgent/response"
	"net/http"

	"StreamAgent/internal/logic"
	"StreamAgent/internal/svc"
	"StreamAgent/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req)
		if err != nil {
			response.Response(w, resp, err)
		} else {
			response.Response(w, resp, nil)
		}
	}
}
