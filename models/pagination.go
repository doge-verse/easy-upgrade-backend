package models

import (
	"github.com/doge-verse/easy-upgrade-backend/api/request"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// Page PageSize Number of data items per page
const (
	PageNum  = 1
	PageSize = 20
)

func DefaultPageNum(pageNum interface{}) uint {
	num := cast.ToUint(pageNum)
	if num > 0 {
		return num
	}
	return PageNum
}

func DefaultPageSize(pageSize interface{}) uint {
	size := cast.ToUint(pageSize)
	if size > 0 {
		return size
	}
	return PageSize
}

// Paginate 分页查询
func Paginate(pageInfo request.PageInfo) func(db *gorm.DB) *gorm.DB {
	if pageInfo.PageNum == 0 {
		pageInfo.PageNum = PageNum
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = PageSize
	}
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageInfo.PageNum - 1) * pageInfo.PageSize
		return db.Offset(int(offset)).Limit(int(pageInfo.PageSize))
	}
}
