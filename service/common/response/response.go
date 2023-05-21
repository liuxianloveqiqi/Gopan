package response

import (
	"golang.org/x/net/context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		httpx.ErrorCtx(context.Background(), w, err)
	} else {
		body.Msg = "Success!"
		body.Data = resp
		httpx.OkJson(w, body)
	}

}
