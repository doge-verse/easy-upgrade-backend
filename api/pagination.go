package api

import (
	"github.com/doge-verse/easy-upgrade-backend/pkg"
	"github.com/doge-verse/easy-upgrade-backend/util"

	"github.com/gin-gonic/gin"
)

// Page PageSize Number of data items per page
const (
	Page     = 1
	PageSize = 12
)

// paginationByGet
func paginationByGet(c *gin.Context) *pkg.Pagination {
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	var p uint = Page
	if page != "" {
		temp, err := util.ParseUint(page)
		if err == nil && temp >= 1 {
			p = temp
		}
	}
	var ps uint = PageSize
	if pageSize != "" {
		temp, err := util.ParseUint(pageSize)
		if err == nil && temp >= 1 {
			ps = temp
		}
	}
	return &pkg.Pagination{
		Total:    0,
		PageSize: ps,
		Current:  p,
		Offset:   (p - 1) * ps,
	}
}
