package main

import (
	"Gopan/app/transfer/rmq/internal/config"
	"Gopan/app/transfer/rmq/internal/service"
	"Gopan/common/errorx"
	"Gopan/common/logs/zapx"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/transfer.yaml", "the etc file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	srv := service.NewService(c)
	queue := kq.MustNewQueue(c.KqConsumerConf, kq.WithHandle(srv.Consume))
	defer queue.Stop()
	// 自定义错误
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)
	fmt.Println("seckill started!!!")
	queue.Start()
}
