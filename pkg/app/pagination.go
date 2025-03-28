package app

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rongpengju/gin-template/configs"
)

type Pagination struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
}

func NewPagination(c *gin.Context) *Pagination {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize <= 0 {
		pageSize = configs.Conf.App.PaginationDefaultSize
	}
	if pageSize > configs.Conf.App.PaginationMaxSize {
		pageSize = configs.Conf.App.PaginationMaxSize
	}

	return &Pagination{Page: page, PageSize: pageSize}
}

func (p *Pagination) GetPage() int {
	return p.Page
}

func (p *Pagination) GetPageSize() int {
	return p.PageSize
}

func (p *Pagination) GetTotalRows() int64 {
	return p.TotalRows
}

func (p *Pagination) SetTotalRows(total int64) {
	p.TotalRows = total
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.PageSize
}
