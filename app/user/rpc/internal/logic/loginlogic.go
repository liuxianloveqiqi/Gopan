package logic

import (
	"Gopan/app/user/model"
	"Gopan/app/user/rpc/internal/svc"
	"Gopan/app/user/rpc/types/user"
	"Gopan/common/errorx"
	"Gopan/common/utils"
	"context"
	"github.com/pkg/errors"
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
	r := l.svcCtx.MasterDb.Where("user_phone = ? or user_email = ?", in.PhoneOrEmail, in.PhoneOrEmail).First(&user0)
	if r.RowsAffected == 0 {
		return nil, errors.Wrapf(errorx.NewCodeError(10005, errorx.ERRPhoneOrEmail), "mobile:%s,phone:%v", in.PhoneOrEmail, in.PhoneOrEmail)
	}
	if r.Error != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(r.Error.Error()), "mobile:%s,phone:%v", in.PhoneOrEmail, in.PhoneOrEmail)

	}
	if !utils.ValidMd5Password(in.PassWord, "liuxian", user0.PassWord) {
		return nil, errors.Wrapf(errorx.NewCodeError(10006, errorx.ERRLoginPassword), "password:%v", in.PassWord)
	}
	return &user.CommonResp{
		UserId: user0.UserId,
	}, nil
}
