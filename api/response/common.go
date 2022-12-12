package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	PageNum  uint        `json:"pageNum"`
	PageSize uint        `json:"pageSize"`
}
