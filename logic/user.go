package logic

import (
	"errors"
	"lelForum/database/postgres"
	"lelForum/models"
	"lelForum/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// See if user exists
	var exist bool
	exist, err = postgres.CheckUserExistence(p.Username)
	if err != nil {
		// DB error
		return err
	}
	if exist {
		// User already exists
		return errors.New("user already exists")
	}

	// Gen user ID
	userID, err := snowflake.GetID()
	if err != nil {
		return err
	}

	// Create user model
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// Save to DB
	return postgres.InsertUser(user)
}
