package init_db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// gorm初始化
func InitGorm(MysqlDataSourece string) *gorm.DB {
	// 将日志写进kafka
	logx.SetWriter(*LogxKafka())
	err := logx.Close()
	if err != nil {
		return nil
	}
	db, err := gorm.Open(mysql.Open(MysqlDataSourece),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				//TablePrefix:   "tech_", // 表名前缀，`User` 的表名应该是 `t_users`
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
		})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接mysql数据库成功")
	}
	return db
}

// redis初始化
func InitRedis() *redis.ClusterClient {
	// Redis集群连接参数
	clusterOptions := &redis.ClusterOptions{
		Addrs: []string{
			"43.139.195.17:6379",
			"43.139.195.17:6380",
			"43.139.195.17:6381",
			"43.139.195.17:6382",
			"43.139.195.17:6383",
			"43.139.195.17:6384",
		},
	}

	// 创建Redis集群客户端
	rdb := redis.NewClusterClient(clusterOptions)

	// 连接Redis集群
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("连接redis失败, error=" + err.Error())
	}
	fmt.Println("redis连接成功")
	return rdb
}
