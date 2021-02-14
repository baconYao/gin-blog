package app

import (
	"github.com/baconYao/gin-blog/global"
	"github.com/baconYao/gin-blog/pkg/convert"
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
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		// DefaultPageSize is 10
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		// MaxPageSize is 100
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
