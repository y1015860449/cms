package service

import (
	"cms/news/conf"
	"github.com/cms/common/name"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	goTracing "github.com/opentracing/opentracing-go"
	"github.com/y1015860449/go-tools/hy_log/hy_zap"
	"time"
)

func NewService(conf *conf.Config) micro.Service {
	// 使用etcd作为默认注册中心
	var etcdRegistry = etcdv3.NewRegistry(func(opt *registry.Options) {
		opt.Addrs = conf.Etcd.Addrs
		//etcdv3.Auth("username", "password")(opt)
	})

	// 创建服务，除了服务名，其它选项可加可不加，比如Version版本号、Metadata元数据等
	srv := micro.NewService(
		micro.Name(name.ServiceNewsSrv),
		micro.Version(conf.Micro.Version),
		micro.RegisterTTL(time.Duration(conf.Micro.RegisterTTL)*time.Second),
		micro.RegisterInterval(time.Duration(conf.Micro.RegisterInterval)*time.Second),
		micro.Registry(etcdRegistry),         // 使用etcd注册中心
		micro.Transport(grpc.NewTransport()), // 默认传输方式模式是http,改为grpc
		micro.WrapHandler(
			// 监控
			prometheus.NewHandlerWrapper(),
			// 链路追踪
			opentracing.NewHandlerWrapper(goTracing.GlobalTracer()),
		),
	)

	hy_zap.ZapLog.Info("init service success")

	return srv
}
