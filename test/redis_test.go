package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"testing"
	"time"
)

func TestRdis(t *testing.T) {
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
		// 如果您的Redis集群有密码，请取消下面这行的注释，并将<password>替换为实际密码
		// Password: "<password>",
	}

	// 创建Redis集群客户端
	client := redis.NewClusterClient(clusterOptions)

	// 连接Redis集群
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := client.Ping(ctx).Err()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 执行 CLUSTER INFO 命令
	info, err := client.ClusterInfo(ctx).Result()
	if err != nil {
		fmt.Println("Failed to execute CLUSTER INFO:", err)
		panic(err)
		return
	}
	// 打印集群信息
	fmt.Println(info)
	fmt.Println("-------------------")
	client.Set(ctx, "19870427401", 5369, 0)

	fmt.Println(client.Get(ctx, "19870427401").Result())
	// 打印集群信息

}
