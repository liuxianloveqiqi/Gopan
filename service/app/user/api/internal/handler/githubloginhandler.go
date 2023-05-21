package handler

import (
	"Gopan/service/app/user/api/internal/logic"
	"Gopan/service/app/user/api/internal/svc"
	"Gopan/service/common/response"
	"fmt"
	"net/http"
)

func GithubLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGithubLoginLogic(r.Context(), svcCtx)
		err := l.GithubLogin()
		authURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", svcCtx.Config.Github.ClientID, svcCtx.Config.Github.RedirectUrl)
		if err != nil {
			response.Response(w, nil, err) //②
			return
		}
		http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
	}
}
