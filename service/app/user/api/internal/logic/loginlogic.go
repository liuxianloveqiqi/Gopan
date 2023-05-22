package logic

import (
	"Gopan/service/app/user/rpc/types/user"
	"Gopan/service/common/errorx"
	"Gopan/service/common/utils"
	"context"
	"github.com/google/uuid"

	"Gopan/service/app/user/api/internal/svc"
	"Gopan/service/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.TokenResp, err error) {
	// todo: add your logic here and delete this line
	cnt, err := l.svcCtx.Rpc.Login(l.ctx, &user.LoginReq{
		PhoneOrEmail: req.PhoneOrEmail,
		PassWord:     req.PassWord,
	})
	if err != nil {
		return nil, err
	}
	accessTokenString, refreshTokenString := utils.GetToken(cnt.UserId, uuid.New().String())
	if accessTokenString == "" || refreshTokenString == "" {
		return nil, errorx.NewCodeError(100002, errorx.JWt)
	}

	return &types.TokenResp{
		UserId:       cnt.UserId,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
