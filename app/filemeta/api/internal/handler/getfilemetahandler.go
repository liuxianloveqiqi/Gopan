package handler

import (
	"Gopan/app/filemeta/api/internal/logic"
	"Gopan/app/filemeta/api/internal/svc"
	"Gopan/app/filemeta/api/internal/types"
	"Gopan/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetFileMetaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFileMetaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetFileMetaLogic(r.Context(), svcCtx)
		resp, err := l.GetFileMeta(&req)
		response.Response(r, w, resp, err) //②

	}
}
