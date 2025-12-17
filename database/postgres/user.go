package postgres

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"lelForum/models"
)

var secret = "xclele"
var (
	ErrorUserExist       = errors.New("user already exists")
	ErrorUserNotExist    = errors.New("user does not exist")
	ErrorInvalidPassword = errors.New("invalid password")
)

// Encapsulate user-related DB operations,
// and provide functions for the logic layer to call

// CheckUserExistence checks if a user with the given username exists
func CheckUserExistence(username string) (err error) {
	sqlStr := `SELECT COUNT(user_id) FROM "user" WHERE username=$1`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func GetUserByID(userID uint64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `SELECT user_id, username FROM "user" WHERE user_id=$1`
	err = db.Get(user, sqlStr, userID)
	return
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

func Login(user *models.User) (err error) {
	originalPasswd := user.Password
	sqlStr := `SELECT user_id,username,password FROM "user" WHERE username=$1`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return
	}
	if encryptPassword(originalPasswd) != user.Password {
		return ErrorInvalidPassword
	}
	return
}
