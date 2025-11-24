package logic

import (
	"lelForum/database/postgres"
	"lelForum/pkg/snowflake"
)

func SignUp() {
	// See if user exists
	postgres.QueryUserByUsername()
	// Gen user ID
	snowflake.GetID()
	// Save to DB
	postgres.InsertUser()
	// redis....
}
