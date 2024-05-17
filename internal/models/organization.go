package models

type Organization struct {
	ID       uint        `json:"id"`
	Name     string      `json:"name"`
	Slug     string      `json:"slug"`
	Metadata interface{} `json:"metadata"`
}
