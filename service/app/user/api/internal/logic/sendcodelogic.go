package logic

import (
	"Gopan/service/app/user/rpc/types/user"
	"Gopan/service/common/errorx"
	"Gopan/service/common/utils"
	"context"
	"fmt"

	"Gopan/service/app/user/api/internal/svc"
	"Gopan/service/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendcodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendcodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendcodeLogic {
	return &SendcodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendcodeLogic) Sendcode(req *types.RegisterByPhoneRep) (resp *types.RegisterByPhoneResp, err error) {
	// todo: add your logic here and delete this line
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errorx.NewCodeError(100001, fmt.Sprintf("sendcode validate校验错误: %v", err))
	}

	cnt, err := l.svcCtx.Rpc.SendCode(l.ctx, &user.SendCodeReq{UserPhone: req.UserPhone})
	if err != nil {
		return nil, err
	}

	return &types.RegisterByPhoneResp{VeCode: cnt.VeCode}, nil
}
