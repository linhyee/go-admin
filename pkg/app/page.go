package app

import (
	"blog/global"
	"blog/pkg/convert"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	ps := convert.StrTo(c.Query("page_size")).MustInt()
	if ps <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if ps > global.AppSetting.DefaultPageSize {
		return global.AppSetting.MaxPageSize
	}
	return ps
}

func GetPageOffset(page, ps int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * ps
	}
	return result
}
