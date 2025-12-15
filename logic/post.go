package logic

import (
	"lelForum/database/postgres"
	"lelForum/models"
	"lelForum/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//generate post ID
	postID, err := snowflake.GetID()
	if err != nil {
		return
	}
	//save to the database
	p.ID = postID
	return postgres.CreatePost(p)
}
