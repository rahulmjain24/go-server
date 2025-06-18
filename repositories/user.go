package repositories

import (
	"database/sql"

	"github.com/rahulmjain24/go-server/db"
	"github.com/rahulmjain24/go-server/models"
)

func CreateUser(name, email string) error {
	_, err := db.DB.Exec("INSERT INTO tbl_users (name, email) VALUES ($1, $2)", name, email)
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.QueryRow("SELECT id, name, email FROM users WHERE email=$1", email).
		Scan(&user.Id, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
