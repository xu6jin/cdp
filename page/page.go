package page

import "fmt"

type Pagination interface {
	Limit() (int, int)
	OrderBy() string
	SetTotalItems(int)
	Extract() map[string]interface{}
}

type pagination struct {
	TotalItems  int    `json:"totalItems"`
	Page        int    `json:"page"`
	RowsPerPage int    `json:"rowsPerPage"`
	SortBy      string `json:"sortBy"`
	Descending  bool   `json:"descending"`
}

func (p *pagination) Limit() (int, int) {
	return p.RowsPerPage, (p.Page - 1) * p.RowsPerPage
}

func (p *pagination) OrderBy() string {
	if p.Descending {
		return fmt.Sprintf("%s DESC", p.SortBy)
	}
	return p.SortBy
}

func (p *pagination) SetTotalItems(totalItems int) {
	p.TotalItems = totalItems
}

func (p *pagination) Extract() map[string]interface{} {
	return map[string]interface{}{
		"totalItems":  p.TotalItems,
		"page":        p.Page,
		"rowsPerPage": p.RowsPerPage,
		"sortBy":      p.SortBy,
		"descending":  p.Descending,
	}
}

func NewDefaultPagination() Pagination {
	return NewPagination(1, 10, "id", true)
}

func NewNoLimitPagination() Pagination {
	return NewPagination(1, -1, "id", true)
}

func NewPagination(page int, rowsPerPage int, sortBy string, descending bool) Pagination {
	if page < 1 {
		page = 1
	}
	if rowsPerPage < 1 {
		rowsPerPage = -1
	}
	if sortBy == "" {
		sortBy = "id"
	}
	return &pagination{Page: page, RowsPerPage: rowsPerPage, SortBy: sortBy, Descending: descending}
}
