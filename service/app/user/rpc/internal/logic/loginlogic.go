package logic

import (
	"Gopan/service/app/user/model"
	"Gopan/service/common/errorx"
	"Gopan/service/common/utils"
	"context"
	"errors"

	"Gopan/service/app/user/rpc/internal/svc"
	"Gopan/service/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.CommonResp, error) {
	// todo: add your logic here and delete this line
	user0 := model.User{}
	r := l.svcCtx.SlaveDb.Where("user_phone = ? or user_email = ?", in.PhoneOrEmail, in.PhoneOrEmail).First(&user0)
	if r.RowsAffected == 0 {
		return nil, errors.New("10005:" + errorx.ERRPhoneOrEmail)
	}
	if r.Error != nil {
		return nil, errorx.NewDefaultError(r.Error.Error())
	}

	if !utils.ValidMd5Password(in.PassWord, "liuxian", user0.PassWord) {
		return nil, errors.New("10006:" + errorx.ERRLoginPassword)
	}
	return &user.CommonResp{
		UserId: user0.UserId,
	}, nil
}
