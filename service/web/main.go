package main

import (
	"context"
	"simpledouyin/service/web/auth"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"simpledouyin/config"
)

func main() {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.WebServiceName),
		provider.WithExportEndpoint(config.EnvConfig.EXPORT_ENDPOINT),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	tracer, cfg := tracing.NewServerTracer()
	h := server.Default(
		server.WithHostPorts(config.WebServiceAddr),
		server.WithMaxRequestBodySize(config.EnvConfig.MAX_REQUEST_BODY_SIZE),
		tracer,
	)
	h.Use(gzip.Gzip(gzip.DefaultCompression))

	h.Use(tracing.ServerMiddleware(cfg))
	pprof.Register(h)

	douyin := h.Group("/douyin")

	// user service
	userGroup := douyin.Group("/user")
	userGroup.POST("/register/", auth.Register)

	h.Spin()
}
