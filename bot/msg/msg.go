package msg

const (
	WellcomeBack = "歡迎回來，%s"
	Wellcome     = "歡迎來到 1li.tw，%s"
	Done         = "完成"
	ShortURL     = "%s/%s -> %s"

	// error
	ServerError                = "伺服器錯誤，請稍後再試"
	Register                   = "請先使用 /start 註冊使用者"
	PermissionDenied           = "權限不足"
	PermissionDeniedCustomCode = "權限不足，你沒有權限自訂短網址"
	InvalidURL                 = "無效的網址"
	URLExist                   = "短網址已存在"

	// perms
	UserNotExist    = "使用者不存在"
	UserNotSelected = "請選擇使用者"
	PermSet         = "設定 %s: %v 給 %s"
	KeyboardClear   = "清除鍵盤"
	PermNotSelected = "請選擇權限"
	DoesSetPerm = "是否要給予權限 %s？"
)
