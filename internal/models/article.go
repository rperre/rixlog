package models

import (
	"errors"
	"rixlog/internal/databases"
)

type ArticleModel struct {
	ID     string `db:"id" json:"id"`
	UserID int64  `db:"user_id" json:"user_id"`
	Title  string `db:"title" json:"title"`
	Body   string `db:"body" json:"body"`
	Slug   string `db:"slug" json:"slug"`
}

func (a *ArticleModel) JSON() *ArticleModel {
	return &ArticleModel{
		ID:     a.ID,
		UserID: a.UserID,
		Title:  a.Title,
		Body:   a.Body,
		Slug:   a.Slug,
	}
}

func (a *ArticleModel) GetByID(id int64) (*ArticleModel, error) {
	Sqlite := databases.Sqlite().Connection
	article := []ArticleModel{}
	if err := Sqlite.Select(&article, "SELECT * FROM article WHERE id=?", id); err != nil {
		return nil, err
	}
	if len(article) == 0 {
		return nil, errors.New("Article not found.")
	}
	return article[0].JSON(), nil
}

func (a *ArticleModel) Create(*ArticleModel) (*ArticleModel, error) { return nil, nil }
func (a *ArticleModel) Edit() (*ArticleModel, error)                { return nil, nil }
func (a *ArticleModel) Delete() (*ArticleModel, error)              { return nil, nil }
