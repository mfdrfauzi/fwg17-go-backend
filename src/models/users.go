package models

import (
	"database/sql"
	"fmt"
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

const limit = 5
const defaultPage = 1

func totalCount(keyword string) (int, error) {
	sql := `
		SELECT COUNT(*) as total
		FROM "users"
		WHERE "email" ILIKE $1
	`
	var total int
	err := db.QueryRow(sql, "%"+keyword+"%").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetAllUser(keyword, sortBy, orderBy string, page int) ([]User, int, error) {
	if page <= 0 {
		page = defaultPage
	}

	column := map[string]bool{"id": true, "email": true, "createdAt": true}
	ordering := map[string]bool{"asc": true, "desc": true}

	if _, ok := column[sortBy]; !ok {
		sortBy = "id"
	}

	if _, ok := ordering[orderBy]; !ok {
		orderBy = "asc"
	}

	offset := (page - 1) * limit

	sql := fmt.Sprintf(`
		SELECT 
		"id", 
		"email",
		"password",
		"createdAt"
		FROM "users"
		WHERE "email" ILIKE $1
		ORDER BY %s %s
		LIMIT %d OFFSET %d
	`, sortBy, orderBy, limit, offset)

	var users []User
	err := db.Select(&users, sql, "%"+keyword+"%")
	if err != nil {
		return nil, 0, err
	}

	totalCount, err := totalCount(keyword)
	if err != nil {
		return nil, 0, err
	}

	totalPage := (totalCount + limit - 1) / limit
	return users, totalPage, nil
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

func UpdateUser(data User) (User, error) {
	sql := `
	UPDATE "users"
	SET
	email=COALESCE(NULLIF(:email,''),email),
	password=COALESCE(NULLIF(:password,''),password)
	WHERE id=:id
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

func DeleteUser(id int) (User, error) {
	sql := `DELETE FROM "users" WHERE id=$1 RETURNING *`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}
