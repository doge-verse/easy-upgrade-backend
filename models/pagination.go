package models

// Pagination generator
type Pagination struct {
	Total    uint `json:"total"`
	PageSize uint `json:"pageSize"`
	Current  uint `json:"current"`
	Offset   uint `json:"-"`
}
