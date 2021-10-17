package initial

import (
	"cms/news/conf"
	"cms/news/dao/cache"
	"cms/news/service"
	"fmt"
	"github.com/cms/common/name"
	"github.com/y1015860449/go-tools/hy_log/hy_zap"
	"github.com/y1015860449/go-tools/hy_servicectrl/hy_tracer"
	"github.com/y1015860449/go-tools/hyredis"
)

func Initial(confPath string) {
	config, err := conf.NewConfig(confPath)
	if err != nil {
		fmt.Printf("init config err (%v)\n", err)
		panic(err)
	}
	// log 初始化
	logConf := hy_zap.DefaultConfig()
	logConf.ServerName = config.Log.ServerName
	logConf.LogPath = config.Log.LogPath
	logConf.MaxAge = config.Log.MaxAge
	logConf.MaxBackups = config.Log.MaxBackups
	logConf.LogLevel = config.Log.LogLevel
	hy_zap.InitLogger(logConf)

	// 链路追踪
	io, err := hy_tracer.InitTracer(name.ServiceNewsSrv, config.Trace.Addr)
	if err != nil {
		hy_zap.ZapLog.Errorf("init tracer err (%v)", err)
	} else {
		defer io.Close()
		hy_zap.ZapLog.Info("init tracer success")
	}

	// 创建 micro service
	srv := service.NewService(config)

	// 初始化缓存
	redisCli, err := hyredis.InitRedis(&hyredis.RedisConfig{
		Addrs:        config.Redis.Addrs,
		Password:     config.Redis.Password,
		MaxIdleConns: config.Redis.MaxIdleConns,
		MaxOpenConns: config.Redis.MaxOpenConns,
		MaxLifeTime:  config.Redis.MaxLifeTime,
	})
	if err != nil {
		hy_zap.ZapLog.Fatalf("init redis err (%v)", err)
	}
	cacheDao := cache.NewCacheOperator(redisCli)
	_ = cacheDao

	// 运行
	if err := srv.Run(); err != nil {
		hy_zap.ZapLog.Fatalf("run server err (%v)", err)
	}
}
