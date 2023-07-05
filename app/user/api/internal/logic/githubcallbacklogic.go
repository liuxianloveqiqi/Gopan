package logic

import (
	"Gopan/app/user/api/internal/svc"
	"Gopan/app/user/api/internal/types"
	"Gopan/app/user/model"
	"Gopan/common/errorx"
	utils "Gopan/common/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GithubCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGithubCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GithubCallbackLogic {
	return &GithubCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GithubCallbackLogic) GithubCallback() (resp *types.TokenResp, err error) {
	// todo: add your logic here and delete this line

	user_auth0 := model.UserAuth{}
	r := l.svcCtx.MysqlDb.Where("provider= ? and provider_id =?", "github", l.ctx.Value("github_id")).First(&user_auth0)
	if r.RowsAffected == 0 {
		fmt.Println("该用户为githubu新用户，开始注册")
		// 新建用户
		user1 := model.User{
			PassWord:   utils.Md5Password(utils.GeneratePassword(10), "liuxian"),
			UserNick:   l.ctx.Value("github_nickname").(string),
			UserSex:    2,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		l.svcCtx.MysqlDb.Create(&user1)
		user1_auth := model.UserAuth{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			UserId:     user1.UserId,
			ProviderId: l.ctx.Value("github_id").(string),
			Provider:   "github",
		}
		l.svcCtx.MysqlDb.Create(&user1_auth)
		accessTokenString, refreshTokenString := utils.GetToken(user1.UserId, uuid.New().String())
		if accessTokenString == "" || refreshTokenString == "" {
			return nil, errorx.NewCodeError(100002, "生成jwt错误")
		}
		return &types.TokenResp{
			UserId:       user1.UserId,
			AccessToken:  accessTokenString,
			RefreshToken: refreshTokenString,
		}, nil
	}

	accessTokenString, refreshTokenString := utils.GetToken(user_auth0.UserId, uuid.New().String())
	if accessTokenString == "" || refreshTokenString == "" {
		return nil, errorx.NewCodeError(100002, "生成jwt错误")

	}
	return &types.TokenResp{
		UserId:       user_auth0.UserId,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
	return
}
