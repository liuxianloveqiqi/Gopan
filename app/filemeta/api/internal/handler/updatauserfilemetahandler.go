package handler

import (
	"Gopan/app/filemeta/api/internal/logic"
	"Gopan/app/filemeta/api/internal/svc"
	"Gopan/app/filemeta/api/internal/types"
	"Gopan/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UpdataUserFileMetaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdataUserFileMetaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdataUserFileMetaLogic(r.Context(), svcCtx)
		err := l.UpdataUserFileMeta(&req)
		response.Response(r, w, nil, err) //②

	}
}
