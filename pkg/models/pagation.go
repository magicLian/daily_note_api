package models

type Pagination struct {
	PageNum  int    `json:"pageNum"`
	Offset   int    `json:"offset"`
	PageSize int    `json:"pageSize"`
	Total    int    `json:"total"`
	Next     string `json:"next"`
}
