package handler

import (
	"Gopan/app/filemeta/api/internal/logic"
	"Gopan/app/filemeta/api/internal/svc"
	"Gopan/common/response"
	"net/http"
)

func GetUserFileMetaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetUserFileMetaLogic(r.Context(), svcCtx)
		resp, err := l.GetUserFileMeta()
		response.Response(r, w, resp, err) //②

	}
}
