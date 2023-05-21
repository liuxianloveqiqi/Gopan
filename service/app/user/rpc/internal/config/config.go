package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MysqlMaster struct {
		DataSource string
	}
	MysqlSlave struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	Credential struct {
		SecretId  string
		SecretKey string
	}
}
