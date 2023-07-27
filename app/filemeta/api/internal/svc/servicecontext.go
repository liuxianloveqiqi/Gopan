package svc

import (
	"Gopan/app/filemeta/api/internal/config"
	"Gopan/app/filemeta/api/internal/middleware"
	"Gopan/app/filemeta/rpc/filemetaclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	JWT    rest.Middleware
	Rpc    filemetaclient.Filemeta
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		JWT:    middleware.NewJWTMiddleware().Handle,
		Rpc:    filemetaclient.NewFilemeta(zrpc.MustNewClient(c.Rpc)),
	}
}
