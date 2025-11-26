package postgres

import (
	"crypto/md5"
	"encoding/hex"
	"lelForum/models"
)

var secret = "xclele"

// Encapsulate user-related DB operations,
// and provide functions for the logic layer to call

// CheckUserExistence checks if a user with the given username exists
func CheckUserExistence(username string) (bool, error) {
	sqlStr := `SELECT COUNT(user_id) FROM "user" WHERE username=$1`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryUserByUsername() {

}

// InsertUser inserts a new user into the database
func InsertUser(user *models.User) (err error) {
	// Encrypt password
	user.Password = encryptPassword(user.Password)
	// Insert user into the database
	sqlStr := `INSERT INTO "user" (user_id, username, password) VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
