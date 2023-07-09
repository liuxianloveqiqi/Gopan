package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MysqlCluster struct {
		DataSource string
	}
	CacheRedis   cache.CacheConf
	RedisCluster struct {
		Cluster1 string
		Cluster2 string
		Cluster3 string
		Cluster4 string
		Cluster5 string
		Cluster6 string
	}
	MinioCluster struct {
		Endpoint  string
		AccessKey string
		SecretKey string
	}
	TecentCOS struct {
		Url       string
		SecretId  string
		SecretKey string
	}
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
}
