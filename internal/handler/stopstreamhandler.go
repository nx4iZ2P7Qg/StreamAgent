package handler

import (
	"StreamAgent/response"
	"net/http"

	"StreamAgent/internal/logic"
	"StreamAgent/internal/svc"
	"StreamAgent/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func StopStreamHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StopStreamReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewStopStreamLogic(r.Context(), ctx)
		resp, err := l.StopStream(req)
		if err != nil {
			response.Response(w, resp, err)
		} else {
			response.Response(w, resp, nil)
		}
	}
}
