package auth_model

// github token
type GithubToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}

// github第三方登录
type GithubConf struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}
