package errno

var (
	ErrorRegisterFail      = NewError(2001001, "注册失败")
	ErrorUserExistFail     = NewError(2001002, "用户已存在")
	ErrorUserNoExistFail   = NewError(2001003, "用户不存在")
	ErrorLoginFail         = NewError(2001004, "登录失败")
	ErrorLogoutFail        = NewError(2001005, "登出失败")
	ErrorPasswordFail      = NewError(2001006, "用户名或密码不正确")
	ErrorCountUserFail     = NewError(2001007, "统计用户失败")
	ErrorGetUserListFail   = NewError(2001008, "获取用户列表失败")
	ErrorUserExist         = NewError(2001009, "用户不存在")
	ErrorUpdateUserFail    = NewError(2001010, "更新用户失败")
	ErrorResetPasswordFail = NewError(2001011, "重置密码失败")
	ErrorOldPasswordFail   = NewError(2001012, "原密码错误")
	ErrorOldNowPswFail     = NewError(2001013, "新密码不能与原密码一致")
	ErrorUserBanFail       = NewError(2001014, "该用户已被禁用")
)
