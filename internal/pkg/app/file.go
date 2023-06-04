package app

import (
	"github.com/DATOULIN/dtservice/internal/pkg/setting"
	"strconv"
)

func BuildSavePath(userId int64, fileName string) string {
	return setting.AppSettings.FileDir + strconv.FormatInt(userId, 10) + "/" + fileName
}
