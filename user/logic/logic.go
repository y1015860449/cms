package logic

import (
	"cms/user/conf"
	"cms/user/dao/cache"
)

type UserLogic struct {
	conf *conf.Config
	cacheDao cache.CacheDao
}