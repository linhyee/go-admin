package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(1000, "服务内部错误")
	InvalidParams             = NewError(1001, "入参错误")
	NotFound                  = NewError(1002, "找不到")
	UnauthorizedAuthNotExists = NewError(1003, "鉴权失败, 找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(1004, "鉴权失败, Token错误")
	UnauthorizedTokenTimout   = NewError(1005, "鉴权失败, Token超时")
	UnauthorizedTokenGenError = NewError(1006, "鉴权失败, Token生成失败")
	TooManyRequest            = NewError(1007, "请求过多")
	ErrorUploadFileFail       = NewError(2001, "上传文件失败")
)
