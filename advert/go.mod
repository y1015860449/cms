module cms/advert

go 1.16

require (
	github.com/cms/common v0.0.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.0.3
	github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/y1015860449/go-tools v0.0.8
	go.uber.org/automaxprocs v1.4.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/cms/common => ../../cms/common
