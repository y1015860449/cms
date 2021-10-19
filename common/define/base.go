package define

type UserStatus int8

const (
	UserVisitor  UserStatus = 0 // 游客
	UserNormal   UserStatus = 1 // 普通用户
	UserIdentify UserStatus = 2 // 认证用户
	UserFreeze   UserStatus = 3 // 冻结用户
	UserWriteOff UserStatus = 4 // 注销用户

)
