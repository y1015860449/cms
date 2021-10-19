package logic

import (
	"cms/user/dao/cache"
	"errors"
	"github.com/cms/common/define"
	"github.com/y1015860449/go-tools/hy_log/hy_zap"
	"github.com/y1015860449/go-tools/hy_utils"
	"time"
)

func (p *UserLogic) MakeToken(userId uint32, userName string) (string, error) {
	if userId <= 0 || len(userName) <= 0 {
		return "", errors.New("param is exception")
	}
	now := time.Now().Unix()
	expiredTime := now + int64(p.conf.Token.ExpiredSec)
	token := hy_utils.GetUUID()
	keyExpiredTime := int64(p.conf.Token.ExpiredSec) + int64(p.conf.Token.SafetySec)
	info := &cache.UserTokenInfo{
		UserId:      userId,
		UserName:    userName,
		ExpiredTime: expiredTime,
		KickOutTime: 0,
	}
	if err := p.cacheDao.SaveUserToken(token, info, keyExpiredTime); err != nil {
		hy_zap.ZapLog.Errorf("cache save user token err(%#v)", err)
		return "", err
	}
	return token, nil
}

func (p *UserLogic) CheckToken(token string, userId uint32, userName string) (define.SvcCode, int64, error) {
	now := time.Now().Unix()
	if len(token) <= 0 || (userId <= 0 && len(userName) <= 0) {
		return define.SvcCode(define.COMMON_PARAM_EXCEPTION), now, errors.New("param is exception")
	}
	info, err := p.cacheDao.GetUserToken(token)
	if err != nil {
		hy_zap.ZapLog.Error("get user token err(%#v)", err)
		return define.SvcCode(define.COMMON_REDIS_ERROR), now, err
	}

	if info.UserId <= 0 && info.UserName == "" && info.ExpiredTime <= 0 {
		return define.SvcCode(define.USER_LOGIN_TOKEN_AUTH_ERROR), now, nil
	}
	if (userId != 0 && userId != info.UserId) || (len(userName) > 0 && userName != info.UserName) {
		return define.SvcCode(define.USER_LOGIN_TOKEN_AUTH_ERROR), now, nil
	}

	if info.KickOutTime != 0 {
		return define.SvcCode(define.USER_LOGIN_TOKEN_KICKOUT_ERROR), info.KickOutTime, nil
	}

	if time.Now().Unix() > info.ExpiredTime {
		return define.SvcCode(define.USER_LOGIN_TOKEN_EXPIRED_ERROR), now, nil
	}

	// 更新redis信息
	info.ExpiredTime = now + int64(p.conf.Token.ExpiredSec)
	keyExpiredTime := int64(p.conf.Token.ExpiredSec) + int64(p.conf.Token.SafetySec)
	if err := p.cacheDao.SaveUserToken(token, info, keyExpiredTime); err != nil {
		hy_zap.ZapLog.Errorf("save user token err(%#v)", err)
		return define.SvcCode(define.COMMON_REDIS_ERROR), now, err
	}
	return define.SvcCode(define.SERVICE_SUCCESS), now, err
}

func (p *UserLogic) KickOutToken(token string, kickTime int64) error {
	if len(token) <= 0 || kickTime <= 0 {
		return errors.New("param is exception")
	}
	info, err := p.cacheDao.GetUserToken(token)
	if err != nil {
		hy_zap.ZapLog.Errorf("get user token(%s) err(%#v)", token, err)
		return err
	}
	if info.UserId <= 0 && info.UserName == "" && info.ExpiredTime <= 0 {
		return nil
	}
	if info.KickOutTime > 0 {
		return nil
	}
	info.KickOutTime = kickTime
	if err := p.cacheDao.SaveUserToken(token, info, int64(p.conf.Token.SafetySec)); err != nil {
		hy_zap.ZapLog.Errorf("save user token err(%#v)", err)
		return err
	}
	return nil
}
