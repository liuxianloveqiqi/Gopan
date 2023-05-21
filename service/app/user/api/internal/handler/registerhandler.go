package handler

import (
	"Gopan/service/app/user/api/internal/logic"
	"Gopan/service/app/user/api/internal/svc"
	"Gopan/service/app/user/api/internal/types"
	"Gopan/service/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		response.Response(w, resp, err) //②

	}
}
