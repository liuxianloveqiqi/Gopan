package handler

import (
	"Gopan/app/filemeta/api/internal/logic"
	"Gopan/app/filemeta/api/internal/svc"
	"Gopan/common/response"
	"net/http"
)

func GetFileSha1Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetFileSha1Logic(r.Context(), svcCtx)
		err := l.GetFileSha1()
		response.Response(r, w, nil, err) //②

	}
}
