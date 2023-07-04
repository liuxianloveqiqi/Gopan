package logic

import (
	"Gopan/app/user/api/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GithubLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGithubLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GithubLoginLogic {
	return &GithubLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GithubLoginLogic) GithubLogin() error {
	// todo: add your logic here and delete this line

	return nil
}
