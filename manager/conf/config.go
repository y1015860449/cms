package conf

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

type Config struct {
	Log     Log   `yaml:"log"`
	Etcd    Etcd  `yaml:"etcd"`
	Kafka   Kafka `yaml:"kafka"`
	Redis   Redis `yaml:"redis"`
	Release bool  `yaml:"release"`
	Trace   Trace `yaml:"trace"`
	Micro   Micro `yaml:"micro"`
}

type Micro struct {
	Version          string `yaml:"version"`
	RegisterTTL      int    `yaml:"registerTTL"`
	RegisterInterval int    `yaml:"registerInterval"`
}

type Trace struct {
	Addr string `yaml:"addr"`
}

type Log struct {
	LogPath    string `yaml:"path"`       // 日志文件路径
	LogLevel   string `yaml:"level"`      //日志级别 debug/info/warn/error
	MaxBackups int    `yaml:"maxbackups"` // 保存的文件个数
	MaxAge     int    `yaml:"maxage"`     // 保存的天数， 没有的话不删除
	ServerName string `yaml:"servername"` // 服务名称
}

type Redis struct {
	Addrs        string `yaml:"addrs"`
	Password     string `yaml:"password"`
	MaxIdleConns int    `yaml:"maxidleconns"` //最初的连接数量
	MaxOpenConns int    `yaml:"maxopenconns"` //连接池最大连接数量,不确定可以用0(0表示自动定义，按需分配)
	MaxLifeTime  int    `yaml:"maxlifetime"`  //连接关闭时间100秒(100秒不使用将关闭)
}

type Etcd struct {
	Addrs []string `yaml:"addrs"`
}

type Kafka struct {
	Addrs []string `yaml:"addrs"`
}

func NewConfig(configFile string) (*Config, error) {
	fileSource := file.NewSource(
		file.WithPath(configFile),
	)
	conf, _ := config.NewConfig()

	if err := conf.Load(fileSource); err != nil {
		return nil, err
	}

	var c Config
	if err := conf.Scan(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
