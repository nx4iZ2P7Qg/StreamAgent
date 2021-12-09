package handler

import (
	"StreamAgent/response"
	"net/http"

	"StreamAgent/internal/logic"
	"StreamAgent/internal/svc"
	"StreamAgent/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func StartStreamHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StartStreamReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewStartStreamLogic(r.Context(), ctx)
		resp, err := l.StartStream(req)
		if err != nil {
			response.Response(w, resp, err)
		} else {
			response.Response(w, resp, nil)
		}
	}
}
