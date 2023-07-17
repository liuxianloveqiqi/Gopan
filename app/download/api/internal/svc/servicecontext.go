package svc

import (
	"Gopan/app/download/api/internal/config"
	"Gopan/app/download/api/internal/middleware"
	"Gopan/common/init_db"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config  config.Config
	JWT     rest.Middleware
	MinioDb *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	minioDb := init_db.InitMinio(c.MinioCluster.Endpoint, c.MinioCluster.AccessKey, c.MinioCluster.SecretKey)
	return &ServiceContext{
		Config:  c,
		JWT:     middleware.NewJWTMiddleware().Handle,
		MinioDb: minioDb,
	}
}
