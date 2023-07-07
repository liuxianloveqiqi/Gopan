package logic

import (
	"Gopan/app/user/model"
	"Gopan/app/user/rpc/internal/svc"
	"Gopan/app/user/rpc/types/user"
	"Gopan/common/errorx"
	"Gopan/common/utils"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"

	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.CommonResp, error) {
	// todo: add your logic here and delete this line
	vc, err := l.svcCtx.Rdb.Get(l.ctx, in.UserPhone).Result()
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(10003, errorx.ERRNoPhone), "该手机号码不存在: %v", in.UserPhone)
	}
	if in.VeCode != vc {
		return nil, errors.Wrapf(errorx.NewCodeError(10004, errorx.ERRValidateCode), "验证码错误：%v", in.VeCode)
	}
	users, err := l.svcCtx.UserModel.FindUserBy(l.svcCtx.MasterDb, "user_phone", in.UserPhone)
	if err != nil {
		return nil, err
	}
	var user0 model.User
	if len(users) == 0 {
		logc.Info(l.ctx, "该用户为新用户，开始注册")
		// 新建用户
		user0 = model.User{
			PassWord:   utils.Md5Password(utils.GeneratePassword(10), "liuxian"),
			UserNick:   utils.RandNickname(),
			UserSex:    2,
			UserPhone:  in.UserPhone,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		l.svcCtx.MasterDb.Create(&user0)
		return &user.CommonResp{
			UserId: user0.UserId,
		}, nil
	} else {
		user0 = users[0]
		logc.Info(l.ctx, "该用户已经注册，直接登陆")
		return &user.CommonResp{
			UserId: user0.UserId,
		}, nil
	}
}
