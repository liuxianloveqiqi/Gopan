package logic

import (
	"Gopan/app/user/api/internal/svc"
	"Gopan/app/user/api/internal/types"
	"Gopan/app/user/rpc/types/user"
	"Gopan/common/errorx"
	"Gopan/common/utils"
	"context"
	"fmt"
	"github.com/pkg/errors"

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
		return nil, errorx.NewCodeError(100001, fmt.Sprintf("validate校验错误: %v", err))
	}

	cnt, err := l.svcCtx.Rpc.SendCode(l.ctx, &user.SendCodeReq{UserPhone: req.UserPhone})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.RegisterByPhoneResp{VeCode: cnt.VeCode}, nil
}
