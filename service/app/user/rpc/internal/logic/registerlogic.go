package logic

import (
	"Gopan/service/app/user/model"
	"Gopan/service/common/errorx"
	"Gopan/service/common/utils"
	"context"
	"fmt"
	"time"

	"Gopan/service/app/user/rpc/internal/svc"
	"Gopan/service/app/user/rpc/types/user"

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
		return nil, errorx.NewCodeError(10003, "redis get错误")
	}
	if in.VeCode != vc {
		return nil, errorx.NewCodeError(10004, "验证码错误")
	}
	users, err := l.svcCtx.UserModel.FindUserBy(l.svcCtx.SlaveDb, "user_phone", in.UserPhone)
	if err != nil {
		return nil, err
	}
	var user0 model.User
	if len(users) == 0 {
		fmt.Println("该用户为新用户，开始注册")
		// 新建用户
		user0 = model.User{
			PassWord:   utils.Md5Password(utils.GeneratePassword(10), "liuxian"),
			UserNick:   utils.RandNickname(),
			UserSex:    2,
			UserPhone:  in.UserPhone,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			DeleteTime: time.Now(),
		}
		l.svcCtx.MasterDb.Create(&user0)
		return &user.CommonResp{
			UserId: user0.UserId,
		}, nil
	} else {
		user0 = users[0]
		fmt.Println("该用户已经注册，直接登陆")
		fmt.Println(user0)
		return &user.CommonResp{
			UserId: user0.UserId,
		}, nil
	}
}
