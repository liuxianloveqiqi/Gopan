// Code generated by goctl. DO NOT EDIT.
// Source: upload.proto

package uploadclient

import (
	"context"

	"Gopan/app/upload/rpc/types/upload"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonResp                 = upload.CommonResp
	FastUploadFileReq          = upload.FastUploadFileReq
	InitialMultipartUploadReq  = upload.InitialMultipartUploadReq
	InitialMultipartUploadResp = upload.InitialMultipartUploadResp
	UploadFileReq              = upload.UploadFileReq
	UploadPartReq              = upload.UploadPartReq

	Upload interface {
		UploadFile(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*CommonResp, error)
		FastUploadFile(ctx context.Context, in *FastUploadFileReq, opts ...grpc.CallOption) (*CommonResp, error)
		InitialMultipartUpload(ctx context.Context, in *InitialMultipartUploadReq, opts ...grpc.CallOption) (*InitialMultipartUploadResp, error)
		UploadPart(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*CommonResp, error)
	}

	defaultUpload struct {
		cli zrpc.Client
	}
)

func NewUpload(cli zrpc.Client) Upload {
	return &defaultUpload{
		cli: cli,
	}
}

func (m *defaultUpload) UploadFile(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := upload.NewUploadClient(m.cli.Conn())
	return client.UploadFile(ctx, in, opts...)
}

func (m *defaultUpload) FastUploadFile(ctx context.Context, in *FastUploadFileReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := upload.NewUploadClient(m.cli.Conn())
	return client.FastUploadFile(ctx, in, opts...)
}

func (m *defaultUpload) InitialMultipartUpload(ctx context.Context, in *InitialMultipartUploadReq, opts ...grpc.CallOption) (*InitialMultipartUploadResp, error) {
	client := upload.NewUploadClient(m.cli.Conn())
	return client.InitialMultipartUpload(ctx, in, opts...)
}

func (m *defaultUpload) UploadPart(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := upload.NewUploadClient(m.cli.Conn())
	return client.UploadPart(ctx, in, opts...)
}
