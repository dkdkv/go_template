package model

import "database/sql"

type PostsModel struct {
	ID       int            `json:"id"`
	Title    string         `json:"title"`
	Content  sql.NullString `json:"content"`
	CoverURL sql.NullString `json:"cover_url"`
}
