package handler

import (
	"Gopan/app/user/api/internal/logic"
	"Gopan/app/user/api/internal/svc"
	"Gopan/common/auth"
	"Gopan/common/auth/auth_model"
	"Gopan/common/response"
	"context"
	"fmt"
	"net/http"
)

func GithubCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从查询参数中获取授权码
		code := r.URL.Query().Get("code")
		// 交换授权码获取访问令牌
		tokenAuthUrl := fmt.Sprintf(
			"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
			svcCtx.Config.Github.ClientID, svcCtx.Config.Github.ClientSecret, code)
		// 获取 token
		var token *auth_model.GithubToken
		var err error
		fmt.Println(888888888888)
		if token, err = auth.GetGithubToken(tokenAuthUrl); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("---------------------------------------%v", token)
		// 通过token，获取用户信息
		user0, err := auth.GetUserInfo(token)
		if err != nil {
			fmt.Println("获取用户信息失败，错误信息为:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "github_id", user0.GithubId))
		r = r.WithContext(context.WithValue(r.Context(), "github_nickname", user0.Nickname))
		l := logic.NewGithubCallbackLogic(r.Context(), svcCtx)
		resp, err := l.GithubCallback()

		response.Response(r, w, resp, err) //②

	}
}
