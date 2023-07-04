package svc

import (
	"Gopan/app/user/model"
	"Gopan/app/user/rpc/internal/config"
	"Gopan/common/init_db"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	Rdb       *redis.ClusterClient
	MasterDb  *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon := sqlx.NewMysql(c.MysqlCluster.DataSource)
	masterDb := init_db.InitGorm(c.MysqlCluster.DataSource)
	masterDb.AutoMigrate(&model.User{}, &model.UserAuth{})
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)
	return &ServiceContext{
		Config:    c,
		MasterDb:  masterDb,
		UserModel: model.NewUserModel(coon, c.CacheRedis),
		Rdb:       redisDb,
	}
}
