package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Rpc         zrpc.RpcClientConf
	MysqlMaster struct {
		DataSource string
	}
	MysqlSlave struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	Github     struct {
		ClientID     string
		RedirectUrl  string
		ClientSecret string
	}
}
