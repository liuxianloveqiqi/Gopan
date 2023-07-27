package svc

import (
	"Gopan/app/filemeta/model"
	"Gopan/app/filemeta/rpc/internal/config"
	"Gopan/common/init_db"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	Rdb     *redis.ClusterClient
	MysqlDb *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := init_db.InitGorm(c.MysqlCluster.DataSource)
	mysqlDb.AutoMigrate(&model.UserFile{}, &model.File{})
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)
	return &ServiceContext{
		Config:  c,
		MysqlDb: mysqlDb,
		Rdb:     redisDb,
	}
}
