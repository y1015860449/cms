package define

type SvcCode int32

const (
	SERVICE_SUCCESS SvcCode = 0
	//common
	COMMON_PARAM_EXCEPTION     SvcCode = 101 //	参数错误
	ERR_UNMARSHAL_FAIL         SvcCode = 102 //	解析失败
	COMMON_MYSQL_ERROR         SvcCode = 103 //	mysql错误
	COMMON_REDIS_ERROR         SvcCode = 104 //  redis错误
	COMMON_MONGODB_ERROR       SvcCode = 105 //	mongodb错误
	COMMON_RPC_SERVICE_ERROR   SvcCode = 106 //	RPC服务错误
	COMMON_ILLEGAL_USER_ERROR  SvcCode = 107 //	非法用户
	COMMON_THIRD_SERVICE_ERROR SvcCode = 108 //  第三方错误

	//用户服务
	USER_LOGIN_TOKEN_AUTH_ERROR    SvcCode = 11001 // token验证错误
	USER_LOGIN_TOKEN_EXPIRED_ERROR SvcCode = 11002 // token过期
	USER_LOGIN_TOKEN_KICKOUT_ERROR SvcCode = 11003 // 在其他设备登陆

)
