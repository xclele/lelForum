package logic

import (
	"lelForum/database/postgres"
	"lelForum/models"
	"lelForum/pkg/snowflake"

	"go.uber.org/zap"
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

func GetPostDetail(pid uint64) (data *models.ApiPostDetail, err error) {
	//get post info from the database
	post, err := postgres.GetPostByID(pid)
	if err != nil {
		zap.L().Error("GetPostID", zap.Error(err))
		return
	}
	//get author name from the database
	user, err := postgres.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("GetUserByID", zap.Error(err))
		return
	}
	//get community info from the database
	community, err := postgres.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("GetCommunityDetail", zap.Error(err))
		return
	}
	//merge the data
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}
