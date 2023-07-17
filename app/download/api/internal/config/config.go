package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MysqlCluster struct {
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
	MinioCluster struct {
		Endpoint  string
		AccessKey string
		SecretKey string
	}
	TencentCOS struct {
		Url       string
		SecretId  string
		SecretKey string
	}
	FileLocalPath string
}
