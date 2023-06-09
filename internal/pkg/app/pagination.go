package app

import (
	"github.com/DATOULIN/dtservice/internal/pkg/convert"
	"github.com/DATOULIN/dtservice/internal/pkg/setting"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page < 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return setting.AppSettings.DefaultPageSize
	}
	if pageSize > setting.AppSettings.MaxPageSize {
		return setting.AppSettings.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		return (page - 1) * pageSize
	}
	return result
}
