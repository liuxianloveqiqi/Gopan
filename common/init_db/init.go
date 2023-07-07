package init_db

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
func InitRedis(redisCluster []string) *redis.ClusterClient {
	// Redis集群连接参数
	clusterOptions := &redis.ClusterOptions{
		Addrs: redisCluster,
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

// minio初始化
func InitMinio(Endpoint, AccessKey, SecretKey string) *minio.Client {
	// 连接到 MinIO 集群
	endpoint := Endpoint
	accessKey := AccessKey
	secretKey := SecretKey
	useSSL := false

	// 创建一个MinIO客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logx.Error("连接minio失败, error=", err)
		panic("连接minio失败, error=" + err.Error())
	}
	fmt.Println("连接minio集群成功")
	return minioClient
}
