package logic

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"Gopan/service/app/user/rpc/internal/svc"
	"Gopan/service/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserList, error) {
	// todo: add your logic here and delete this line

	user0, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.New("10001:" + err.Error())
	}
	users := make([]*user.User, 0)
	user1 := &user.User{
		UserId:     user0.UserId,
		PassWord:   user0.PassWord,
		User_Nick:  user0.UserNick,
		User_Face:  user0.UserFace,
		User_Sex:   user0.UserSex,
		User_Email: user0.UserEmail,
		User_Phone: user0.UserPhone,
		CreateTime: timestamppb.New(user0.CreateTime),
		UpdateTime: timestamppb.New(user0.UpdateTime),
		DeleteTime: timestamppb.New(user0.DeleteTime),
	}
	users = append(users, user1)
	return &user.UserList{
		Users: users,
	}, nil
}
