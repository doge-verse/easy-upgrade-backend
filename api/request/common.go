package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	PageNum  uint `json:"pageNum" form:"pageNum"`
	PageSize uint `json:"pageSize" form:"pageSize"`
}
