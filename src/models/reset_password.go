package models

import (
	"database/sql"
	"time"
)

type ResetForm struct {
	Id        int          `db:"id"`
	Email     string       `db:"email" form:"email"`
	Otp       string       `db:"otp" form:"otp"`
	CreatedAt time.Time    `db:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt"`
}

func ResetByOtp(otp string) (ResetForm, error) {
	sql := `SELECT * FROM "resetPassword" WHERE otp=$1`
	data := ResetForm{}
	err := db.Get(&data, sql, otp)
	return data, err
}

func CreateResetPassword(data ResetForm) (ResetForm, error) {
	sql := `
	INSERT INTO "resetPassword" ("email", "otp")
	VALUES (:email, :otp)
	RETURNING *`

	result := ResetForm{}
	rows, err := db.NamedQuery(sql, &data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, nil
}

func DeleteResetPassword(id int) (User, error) {
	sql := `DELETE FROM "resetPassword" WHERE id=$1 RETURNING *`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}
