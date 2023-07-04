package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Rpc          zrpc.RpcClientConf
	MysqlCluster struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	Github     struct {
		ClientID     string
		RedirectUrl  string
		ClientSecret string
	}
	RedisCluster struct {
		Cluster1 string
		Cluster2 string
		Cluster3 string
		Cluster4 string
		Cluster5 string
		Cluster6 string
	}
}
