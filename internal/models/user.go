package models

import (
	"rixlog/internal/databases"
)

type UserDB struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Name     string `db:"name"`
	Admin    bool   `db:"admin"`
}

type UserJSON struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Admin    bool   `json:"admin"`
}

func (u *UserDB) JSON() *UserJSON {
	return &UserJSON{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Name:     u.Name,
		Admin:    u.Admin,
	}
}

func (u *UserDB) GetByID(id int64) (*UserJSON, error) {
	Sqlite := databases.Sqlite().Connection
	user := []UserDB{}
	if err := Sqlite.Select(&user, "SELECT * FROM user WHERE id=?", id); err != nil {
		return nil, err
	}

	return user[0].JSON(), nil
}

func (u *UserDB) Create(*UserDB) (*UserJSON, error) { return nil, nil }
func (u *UserDB) Edit() (*UserJSON, error)          { return nil, nil }
func (u *UserDB) Delete() (*UserJSON, error)        { return nil, nil }

var _UserModel *UserDB

func User() *UserDB {
	if _UserModel != nil {
		return _UserModel
	}
	_UserModel = &UserDB{}
	return _UserModel
}
