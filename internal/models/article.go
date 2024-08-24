package models

import (
	"rixlog/internal/databases"
)

type ArticleDB struct {
	ID     string `db:"id"`
	UserID int64  `db:"user_id"`
	Title  string `db:"title"`
	Body   string `db:"body"`
	Slug   string `db:"slug"`
}

type ArticleJSON struct {
	ID     string `json:"id"`
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Slug   string `json:"slug"`
}

func (a *ArticleDB) JSON() *ArticleJSON {
	return &ArticleJSON{
		ID:     a.ID,
		UserID: a.UserID,
		Title:  a.Title,
		Body:   a.Body,
		Slug:   a.Slug,
	}
}

func (a *ArticleDB) GetByID(id int64) (*ArticleJSON, error) {
	Sqlite := databases.Sqlite().Connection
	article := []ArticleDB{}
	if err := Sqlite.Select(&article, "SELECT * FROM article WHERE id=?", id); err != nil {
		return nil, err
	}

	return article[0].JSON(), nil
}

func (a *ArticleDB) Create(*ArticleDB) (*ArticleJSON, error) { return nil, nil }
func (a *ArticleDB) Edit() (*ArticleJSON, error)             { return nil, nil }
func (a *ArticleDB) Delete() (*ArticleJSON, error)           { return nil, nil }

var _ArticleModel *ArticleDB

func Article() *ArticleDB {
	if _ArticleModel != nil {
		return _ArticleModel
	}

	_ArticleModel = &ArticleDB{}
	return _ArticleModel
}
