package main

import (
	"log"
	"net"
	"simpledouyin/config"
	favoriteService "simpledouyin/kitex_gen/douyin/favorite/favoriteservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	var err error

	r, err := consul.NewConsulRegister(config.EnvConfig.CONSUL_ADDR)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.IP+config.FavoriteServiceAddr)
	if err != nil {
		panic(err)
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.FavoriteServiceName),
		provider.WithExportEndpoint(config.EnvConfig.EXPORT_ENDPOINT),
		provider.WithInsecure(),
	)

	srv := favoriteService.NewServer(
		new(FavoriteServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.FavoriteServiceName,
		}),
	)
	err = srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
