package handler

import (
	"Gopan/app/download/api/internal/logic"
	"Gopan/app/download/api/internal/svc"
	"Gopan/app/download/api/internal/types"
	"Gopan/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func DownloadMinioHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadMinioReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDownloadMinioLogic(r.Context(), svcCtx)
		err := l.DownloadMinio(&req, w, r)
		response.Response(r, w, nil, err) //②

	}
}
