package model

import "database/sql"

func (m *PostsModel) TableName() string {
	return "posts"
}

func (m *PostsModel) Columns() []string {
	return []string{"id", "title", "content", "cover_url"}
}

func (m *PostsModel) Values() []interface{} {
	return []interface{}{m.ID, m.Title, m.Content, m.CoverURL}
}

func (m *PostsModel) Scan(row *sql.Row) error {
	return row.Scan(&m.ID, &m.Title, &m.Content, &m.CoverURL)
}

func (m *PostsModel) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&m.ID, &m.Title, &m.Content, &m.CoverURL)
}
