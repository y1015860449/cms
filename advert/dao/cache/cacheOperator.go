package cache

import (
	"github.com/y1015860449/go-tools/hyredis"
)

type cacheOperator struct {
	redisCli *hyredis.HyRedis
}

func NewCacheOperator(redisCli *hyredis.HyRedis) CacheDao {
	return &cacheOperator{redisCli: redisCli}
}