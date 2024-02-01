package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mfdrfauzi/fwg17-go-backend/src/lib"
)

var db *sqlx.DB = lib.DB

type User struct {
	Id        int          `db:"id" json:"id"`
	Email     string       `db:"email" json:"email" form:"email"`
	Password  string       `db:"password" json:"password" form:"password"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

func GetAllUser() ([]User, error) {
	sql := `SELECT * FROM "users"`
	data := []User{}
	err := db.Select(&data, sql)
	return data, err
}

func FindOneUser(id int) (User, error) {
	sql := `SELECT * FROM "users" WHERE id=$1`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateUser(data User) (User, error) {
	sql := `
	INSERT INTO "users" ("email", "password")
	VALUES (:email, :password)
	RETURNING *`

	result := User{}
	rows, err := db.NamedQuery(sql, &data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, nil
}
