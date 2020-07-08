package models

import (
	"database/sql"
	"fmt"
	"time"
)

// Blog ...
type Blog struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	Summary   string    `json:"summary"`
	Path      string    `json:"path"`
	Date      time.Time `json:"date"`
}

// Save ...
func (m *Blog) Save(db *sql.DB) error {
	request := fmt.Sprintf(`INSERT %s (title, cover, summary, path, date) values (?,?,?,?,?)`, tableBlog)
	_, err := db.Exec(request, m.Title, m.Cover, m.Summary, m.Path, m.Date)
	return err
}

// Update into database
func (m *Blog) Update(db *sql.DB) error {
	return nil
}

// Load Blog into database
func (m *Blog) Load(db *sql.DB) error {
	request := fmt.Sprintf(`SELECT * FROM %s WHERE id=?`, tableBlog)
	rows, err := db.Query(request, m.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var notEmpty bool
	for rows.Next() {
		notEmpty = true
		if err := rows.Scan(m.toScan()...); err != nil {
			return err
		}
	}
	if !notEmpty {
		return fmt.Errorf("empty set")
	}
	return nil
}

// Digest ...
func (m *Blog) Digest() interface{} {
	return fmt.Sprintf("blog-%s-%s", m.Title, m.Path)
}

func (m *Blog) toScan() []interface{} {
	return []interface{}{
		&m.ID, &m.CreatedAt, &m.UpdatedAt, &m.Title, &m.Cover, &m.Summary, &m.Path, &m.Date,
	}
}

// ListBlogs ...
func ListBlogs(db *sql.DB) ([]*Blog, error) {
	request := fmt.Sprintf("SELECT * FROM %s ORDER BY created_at DESC", tableBlog)
	rows, err := db.Query(request)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r := make([]*Blog, 0)
	for rows.Next() {
		var s Blog
		if err := rows.Scan(s.toScan()...); err != nil {
			continue
		}

		r = append(r, &s)
	}
	return r, nil
}
