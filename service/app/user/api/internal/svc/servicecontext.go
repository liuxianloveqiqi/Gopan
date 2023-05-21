package svc

import (
	"Gopan/service/app/user/api/internal/config"
	"Gopan/service/app/user/api/internal/middleware"
	"Gopan/service/app/user/model"
	"Gopan/service/app/user/rpc/userclient"
	"Gopan/service/common/init_db"
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
	MasterDb      *gorm.DB
	SlaveDb       *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon := sqlx.NewMysql(c.MysqlMaster.DataSource)
	mysqlDb := init_db.InitGorm(c.MysqlMaster.DataSource)
	slaveDb := init_db.InitGorm(c.MysqlSlave.DataSource)
	mysqlDb.AutoMigrate(&model.User{}, &model.UserAuth{})
	redisDb := init_db.InitRedis()
	return &ServiceContext{
		Config:        c,
		JWT:           middleware.NewJWTMiddleware().Handle,
		Rpc:           userclient.NewUser(zrpc.MustNewClient(c.Rpc)),
		MasterDb:      mysqlDb,
		SlaveDb:       slaveDb,
		UserModel:     model.NewUserModel(coon, c.CacheRedis),
		UserAuthModel: model.NewUserAuthModel(coon, c.CacheRedis),
		Rdb:           redisDb,
	}
}
