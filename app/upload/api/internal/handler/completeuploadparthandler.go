package handler

import (
	"Gopan/app/upload/api/internal/logic"
	"Gopan/app/upload/api/internal/svc"
	"Gopan/app/upload/api/internal/types"
	"Gopan/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CompleteUploadPartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CompleteUploadPartReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCompleteUploadPartLogic(r.Context(), svcCtx)
		err := l.CompleteUploadPart(&req)
		response.Response(r, w, nil, err) //②

	}
}
