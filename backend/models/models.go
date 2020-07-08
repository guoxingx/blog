package models

import (
	"database/sql"
	"fmt"
)

// Model interface
type Model interface {
	Save(db *sql.DB) error

	Load(db *sql.DB) error

	Update(db *sql.DB) error

	Digest() interface{}
}

const (
	tableBlog = "blog"
)

// Error
var (
	ErrEmptySet = fmt.Errorf("empty set")
)
