package errno

var (
	Success                  = NewError(0, "成功")
	ServerError              = NewError(10000000, "服务内部错误")
	InvalidParams            = NewError(10000001, "入参错误")
	NotFound                 = NewError(10000002, "找不到")
	UnauthorizedNoTokenError = NewError(10000003, "鉴权失败，Token为空")
	UnauthorizedTokenTimeout = NewError(10000004, "鉴权失败，Token过期")
	UnauthorizedTokenError   = NewError(10000005, "鉴权失败，Token不合法")
	TooManyRequests          = NewError(10000006, "请求过多")
)
