package logic

import (
	"Gopan/app/user/api/internal/svc"
	"Gopan/app/user/api/internal/types"
	"Gopan/app/user/rpc/types/user"
	"Gopan/common/errorx"
	"Gopan/common/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.TokenResp, err error) {
	// todo: add your logic here and delete this line
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(100001, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}
	cnt, err := l.svcCtx.Rpc.Register(l.ctx, &user.RegisterReq{
		UserPhone: req.UserPhone,
		VeCode:    req.VeCode,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	accessTokenString, refreshTokenString := utils.GetToken(cnt.UserId, uuid.New().String())
	if accessTokenString == "" || refreshTokenString == "" {
		return nil, errors.Wrapf(errorx.NewCodeError(100002, errorx.JWt), "生成jwt错误 err:%v", err)
	}

	return &types.TokenResp{
		UserId:       cnt.UserId,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
