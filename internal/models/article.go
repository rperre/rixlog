package models

import (
	"errors"
	"net/http"
	"rixlog/internal/databases"
)

type Article struct {
	ID     string `db:"id" json:"id"`
	UserID int64  `db:"user_id" json:"user_id"`
	Title  string `db:"title" json:"title"`
	Body   string `db:"body" json:"body"`
	Slug   string `db:"slug" json:"slug"`
}

func (a *Article) GetByID(id int64) (*Article, error) {
	Sqlite := databases.Sqlite().Connection
	article := []Article{}
	if err := Sqlite.Select(&article, "SELECT * FROM article WHERE id=?", id); err != nil {
		return nil, err
	}
	if len(article) == 0 {
		return nil, errors.New("Article not found.")
	}
	return &article[0], nil
}

func (rd *Article) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *Article) Create(*Article) (*Article, error) { return nil, nil }
func (a *Article) Edit() (*Article, error)           { return nil, nil }
func (a *Article) Delete() (*Article, error)         { return nil, nil }
