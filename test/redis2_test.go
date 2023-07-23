package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestF(t *testing.T) {
	// 创建 Redis 集群客户端实例
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"10.17.232.216:6390", // 节点1的主机和端口
			"10.17.232.216:6391", // 节点2的主机和端口
			"10.17.232.216:6392", // 节点3的主机和端口

		},
	})

	// 测试连接
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}

	fmt.Println("连接成功:", pong)

	client.Set(context.TODO(), "aaaa", "eeeee", 0)
	fmt.Println(client.Get(context.TODO(), "aaaa").Result())
	// 关闭客户端连接
	err = client.Close()
	if err != nil {
		fmt.Println("关闭连接失败:", err)
		return
	}
}
