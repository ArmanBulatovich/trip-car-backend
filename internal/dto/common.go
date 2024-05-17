package dto

type Pagination struct {
	Page    int `form:"page"`
	PerPage int `form:"per_page"`
}
