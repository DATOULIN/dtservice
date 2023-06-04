package errno

var (
	ErrorUploadFail         = NewError(3001001, "文件上传失败")
	UploadSuccess           = NewError(3001002, "上传成功")
	ErrorFileIsRequiredFail = NewError(3001003, "文件不能为空")
)
