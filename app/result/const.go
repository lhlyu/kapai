package result

const (
	ErrUnknown       = "未知异常"
	ErrUserNotfound  = "用户不存在"
	ErrUserPassword  = "密码不正确"
	ErrCreateUser    = "创建用户账号异常"
	ErrAccountCheck  = "账号仅支持字母数字下划线2到12位的字符组成"
	ErrAccountEmpty  = "账号不能为空"
	ErrPasswordEmpty = "密码不能为空"
	ErrToken         = "token生成失败"

	ErrRoomNotfound = "房间不存在"
)
