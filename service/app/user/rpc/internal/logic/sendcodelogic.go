package logic

import (
	"Gopan/service/app/user/rpc/internal/svc"
	"Gopan/service/app/user/rpc/types/user"
	"Gopan/service/common/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCodeLogic) SendCode(in *user.SendCodeReq) (*user.SendCodeResp, error) {
	// todo: add your logic here and delete this line

	vecode := utils.SMS(in.UserPhone, l.svcCtx.Config.Credential.SecretId, l.svcCtx.Config.Credential.SecretKey, l.ctx, l.svcCtx.Rdb)
	return &user.SendCodeResp{
		VeCode: vecode,
	}, nil
}
