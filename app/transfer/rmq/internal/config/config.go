package config

import (
	"github.com/zeromicro/go-queue/kq"
)

type Config struct {
	KqConsumerConf kq.KqConf
	MysqlCluster   struct {
		DataSource string
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
