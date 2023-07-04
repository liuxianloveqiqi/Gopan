package auth

import (
	"Gopan/common/auth/auth_model"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type GithubUser struct {
	GithubId string
	Nickname string
}

// 获取用户登录的请求
func GetGithubApiRequest(url string, token *auth_model.GithubToken) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))
	return req, nil
}

// 获取用户信息
func GetUserInfo(token *auth_model.GithubToken) (GithubUser, error) {

	// 使用github提供的接口
	var userInfoUrl = "https://api.github.com/user"
	req, err := GetGithubApiRequest(userInfoUrl, token)
	if err != nil {
		return GithubUser{}, err
	}

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return GithubUser{}, err
	}
	// 将响应的数据写入userInfo中，并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		fmt.Println(err)
		return GithubUser{}, err
	}
	fmt.Println(reflect.TypeOf(userInfo["name"].(string)).Kind())
	fmt.Println(reflect.TypeOf(userInfo["id"]).Kind())
	fmt.Println(userInfo["name"].(string), "  这里是name")

	user := GithubUser{
		Nickname: userInfo["name"].(string),
		GithubId: fmt.Sprintf("%.6f", userInfo["id"].(float64)),
	}
	fmt.Println("结构体user", user)
	return user, nil
}

// 获取 github token
func GetGithubToken(url string) (*auth_model.GithubToken, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var token auth_model.GithubToken
	if err := json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}
