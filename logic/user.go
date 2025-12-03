package logic

import (
	"lelForum/database/postgres"
	"lelForum/models"
	"lelForum/pkg/jwt"
	"lelForum/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// See if user exists
	err = postgres.CheckUserExistence(p.Username)
	if err != nil {
		// Exists or other error
		return
	}
	// Gen user ID
	userID, err := snowflake.GetID()
	if err != nil {
		return
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

func Login(p *models.ParamLogin) (string, error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// Here user is a pointer, and the userid will be filled in postgres.Login
	if err := postgres.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
}
