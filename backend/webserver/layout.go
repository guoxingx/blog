package webserver

import (
	"fmt"
	"time"

	"backend/models"
)

type blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Cover   string `json:"cover"`
	Summary string `json:"summary"`
	Path    string `json:"path"`
	Date    string `json:"date"`
}

func (b *blog) fromModel(m *models.Blog) *blog {
	b.ID = m.ID
	b.Title = m.Title
	b.Cover = m.Cover
	b.Summary = m.Summary
	b.Path = m.Path
	b.Date = m.Date.Format("2006-01-02 15:04:05")
	return b
}

func (b *blog) toModel() (*models.Blog, error) {
	date, err := time.Parse("2006-01-02 15:04:05", b.Date)
	if err != nil {
		return nil, err
	}
	if date.After(time.Now()) || date.Before(time.Now().Add(-time.Hour*24*365*5)) {
		return nil, fmt.Errorf("error date: %v", date)
	}

	m := models.Blog{
		Title:   b.Title,
		Cover:   b.Cover,
		Summary: b.Summary,
		Path:    b.Path,
		Date:    date,
	}
	return &m, nil
}
