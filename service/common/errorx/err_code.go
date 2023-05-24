package errorx

const (
	JWt              = "生成jwt错误"                //10002
	ERRNoPhone       = "该手机号码不存在"               //10003
	ERRValidateCode  = "验证码错误"                  //10004
	ERRPhoneOrEmail  = "手机号或者邮箱错误"              //10005
	ERRLoginPassword = "登陆密码错误"                 // 10006
	ERRStatusError   = "无法从 status.Error 中提取信息" //-1
)

const (
	ErrMysqlDateNoResult = "数据库查询错误" //20001
)

const (
	ErrHeadNil    = "请求头中auth为空"   //30001
	ErrHeadFormat = "请求头中auth格式有误" //30002
	ErrTokenProve = "token验证错误"    // 30003
)
