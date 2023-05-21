package svc

import (
	"Gopan/service/app/user/model"
	"Gopan/service/app/user/rpc/internal/config"
	"Gopan/service/common/init_db"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	Rdb       *redis.ClusterClient
	MasterDb  *gorm.DB
	SlaveDb   *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon := sqlx.NewMysql(c.MysqlMaster.DataSource)
	mysqlDb := init_db.InitGorm(c.MysqlMaster.DataSource)
	slaveDb := init_db.InitGorm(c.MysqlSlave.DataSource)
	mysqlDb.AutoMigrate(&model.User{}, &model.UserAuth{})
	redisDb := init_db.InitRedis()
	return &ServiceContext{
		Config:    c,
		MasterDb:  mysqlDb,
		SlaveDb:   slaveDb,
		UserModel: model.NewUserModel(coon, c.CacheRedis),
		Rdb:       redisDb,
	}
}
