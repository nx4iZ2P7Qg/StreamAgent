package response

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

type Body struct {
	Code    int         `json:"code"`
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = -1
		body.Message = err.Error()
		body.Type = "error"
	} else {
		body.Message = "ok"
		body.Result = resp
		body.Type = "success"
	}
	httpx.OkJson(w, body)
}
