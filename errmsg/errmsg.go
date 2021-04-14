package errmsg


const (
	Success = 200
	Error = 500

	//token类错误
	InvalidToken      = 1001
	TokenNotExist     = 1002
	TokenError        = 1003
	AuthEmpty         = 1004
	TokenRunTimeError = 1005

	//请求类错误
	ErrRequest = 2001

	//查询数据库类错误
	ErrInfoNotFound = 3001
	ErrDatabaseFound = 3002

	//用户类存在
	ErrUserNameUsed = 4001
	ErrUserEmailUsed = 4002
	ErrUserPhoneUsed = 4003
	ErrPassword = 4004
	ErrPhoneNotExist = 4005
	ErrPasswordDifferent = 4006
	ErrEmailNotExist = 4007

	ErrEmailCode = 5001
)

var CodeMsg = map[int]string{

	Success: "成功",
	Error: "失败",

	InvalidToken: "非法的token",
	TokenError: "token错误",
	AuthEmpty: "请求头中的auth为空",
	TokenNotExist: "token不存在",
	TokenRunTimeError: "token过期",


	ErrRequest: "请求错误",

	ErrInfoNotFound: "未查找到相关信息",
	ErrDatabaseFound: "数据库查找错误",

	ErrUserNameUsed: "用户名已存在",
	ErrUserEmailUsed: "用户邮箱已存在",
	ErrUserPhoneUsed: "用户电话已存在",
	ErrPassword: "用户密码错误",
	ErrPhoneNotExist: "号码不存在",
	ErrEmailNotExist: "邮箱不存在",
	ErrPasswordDifferent: "密码不一致",

	ErrEmailCode: "邮箱验证码错误",
}