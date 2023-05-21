package handler

import (
	"Gopan/service/app/user/api/internal/logic"
	"Gopan/service/app/user/api/internal/svc"
	"Gopan/service/app/user/api/internal/types"
	"Gopan/service/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err) //②

	}
}
