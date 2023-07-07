package main

import (
	"Gopan/common/errorx"
	"Gopan/common/logs/zapx"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"Gopan/app/upload/api/internal/config"
	"Gopan/app/upload/api/internal/handler"
	"Gopan/app/upload/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/upload-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
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
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
