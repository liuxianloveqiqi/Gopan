// Code generated by goctl. DO NOT EDIT.
// Source: upload.proto

package server

import (
	"context"

	"Gopan/app/upload/rpc/internal/logic"
	"Gopan/app/upload/rpc/internal/svc"
	"Gopan/app/upload/rpc/types/upload"
)

type UploadServer struct {
	svcCtx *svc.ServiceContext
	upload.UnimplementedUploadServer
}

func NewUploadServer(svcCtx *svc.ServiceContext) *UploadServer {
	return &UploadServer{
		svcCtx: svcCtx,
	}
}

func (s *UploadServer) UploadFile(ctx context.Context, in *upload.UploadFileReq) (*upload.CommonResp, error) {
	l := logic.NewUploadFileLogic(ctx, s.svcCtx)
	return l.UploadFile(in)
}
