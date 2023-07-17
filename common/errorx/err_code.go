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
	ErrMysqlDateNoResult = "用户信息查询为空"         //20001
	ErrFileCreat         = "file表Creat错误"     //20002
	ErrUserFlileCreat    = "userfile表Creat错误" //20003
)

const (
	ErrHeadNil    = "请求头中auth为空"   //30001
	ErrHeadFormat = "请求头中auth格式有误" //30002
	ErrTokenProve = "token验证错误"    // 30003
)

const (
	ErrFileOpen                  = "打开文件错误"                //40001
	ErrKafkaFileMeta             = "kafka异步fileMeta失败"     //40002
	ErrKafkaUserFileMeta         = "kafka异步userfileMate失败" //40003
	ErrMultipartUploadNoComplete = "分块上传的文件不完整"            // 40004
	ErrFileSha1Falsify           = "文件sha1值校验错误，文件已被篡改"    //40004
)
