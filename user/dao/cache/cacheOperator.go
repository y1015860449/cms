package cache

import (
	"encoding/json"
	"fmt"
	"github.com/y1015860449/go-tools/hyredis"
)

type cacheOperator struct {
	redisCli *hyredis.HyRedis
}

var preTokenKey = "cms:user:token:%s"

func getTokenKey(token string) string {
	return fmt.Sprintf(preTokenKey, token)
}

func (p *cacheOperator) SaveUserToken(loginToken string, info *UserTokenInfo, expired int64) error {
	key := getTokenKey(loginToken)
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return p.redisCli.SetEx(key, string(data), int(expired))
}

func (p *cacheOperator) GetUserToken(loginToken string) (*UserTokenInfo, error) {
	key := getTokenKey(loginToken)
	rest, err := p.redisCli.Get(key)
	if err != nil {
		return nil, err
	}
	var info UserTokenInfo
	if len(rest) > 0 {
		if err := json.Unmarshal([]byte(rest), &info); err != nil {
			return &info, err
		}
	}
	return &info, nil
}

func (p *cacheOperator) DeleteUserToken(loginToken string) error {
	key := getTokenKey(loginToken)
	err := p.redisCli.Del(key)
	return err
}


func NewCacheOperator(redisCli *hyredis.HyRedis) CacheDao {
	return &cacheOperator{redisCli: redisCli}
}