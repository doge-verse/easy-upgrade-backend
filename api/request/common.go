package request

import "github.com/golang-jwt/jwt"

// PageInfo Paging common input parameter structure
type PageInfo struct {
	PageNum  uint `json:"pageNum" form:"pageNum"`
	PageSize uint `json:"pageSize" form:"pageSize"`
}

// CustomClaims .
type CustomClaims struct {
	UserID uint
	jwt.StandardClaims
}
