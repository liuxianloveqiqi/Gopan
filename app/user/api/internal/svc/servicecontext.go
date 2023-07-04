package svc

import (
	"Gopan/app/user/api/internal/config"
	"Gopan/app/user/api/internal/middleware"
	"Gopan/app/user/model"
	"Gopan/app/user/rpc/userclient"
	"Gopan/common/init_db"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	JWT           rest.Middleware
	Rpc           userclient.User
	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
	Rdb           *redis.ClusterClient
	MysqlDb       *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon := sqlx.NewMysql(c.MysqlCluster.DataSource)
	masterDb := init_db.InitGorm(c.MysqlCluster.DataSource)
	masterDb.AutoMigrate(&model.User{}, &model.UserAuth{})
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)
	return &ServiceContext{
		Config:        c,
		JWT:           middleware.NewJWTMiddleware().Handle,
		Rpc:           userclient.NewUser(zrpc.MustNewClient(c.Rpc)),
		MysqlDb:       masterDb,
		UserModel:     model.NewUserModel(coon, c.CacheRedis),
		UserAuthModel: model.NewUserAuthModel(coon, c.CacheRedis),
		Rdb:           redisDb,
	}
}
