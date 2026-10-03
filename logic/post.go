package logic

import (
	"context"
	"fmt"
	"lelForum/database/postgres"
	"lelForum/database/redis"
	"lelForum/models"
	"lelForum/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(ctx context.Context, p *models.Post) (err error) {
	//generate post ID
	postID, err := snowflake.GetID()
	if err != nil {
		return
	}
	//save to the database
	p.ID = postID
	err = postgres.CreatePost(p)
	if err != nil {
		return
	}
	err = redis.CreatePost(ctx, p.ID, p.CommunityID)
	return
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

// Integrated function to get post list
func GetPostList(ctx context.Context, p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	if p.CommunityID == 0 {
		data, err = GetPostListByCategory(ctx, p)
	} else {
		//get posts by community
		data, err = GetPostListByCommunity(ctx, p)
	}
	if err != nil {
		zap.L().Error("GetPostList", zap.Error(err))
		return nil, err
	}
	return
}

func GetPostListByCategory(ctx context.Context, p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//get post id list from redis
	ids, err := redis.GetPostIDsInOrder(ctx, p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("GetPostListInOrder success, but no result found")
		return
	}
	//get post list from the database based on the post ids
	posts, err := postgres.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	// Get the vote data from Redis
	voteData, err := redis.GetPostVoteData(ctx, ids)
	if err != nil {
		return
	}

	// Create a map from post ID to vote count for correct mapping
	voteMap := make(map[uint64]int64, len(ids))
	for i, idStr := range ids {
		var id uint64
		if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil && i < len(voteData) {
			voteMap[id] = voteData[i]
		}
	}

	//traverse each post to get author and community info
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		//get author name from the database
		user, err := postgres.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("GetUserByID", zap.Error(err))
			continue
		}
		//get community info from the database
		community, err := postgres.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("GetCommunityDetail", zap.Error(err))
			continue
		}
		//merge the data
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteMap[post.ID],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostListByCommunity(ctx context.Context, p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetCommunityPostIDsInOrder(ctx, p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("GetPostListInOrder success, but no result found")
		return
	}
	//get post list from the database based on the post ids
	posts, err := postgres.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	// Get the vote data from Redis
	voteData, err := redis.GetPostVoteData(ctx, ids)
	if err != nil {
		return
	}

	// Create a map from post ID to vote count for correct mapping
	voteMap := make(map[uint64]int64, len(ids))
	for i, idStr := range ids {
		var id uint64
		if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil && i < len(voteData) {
			voteMap[id] = voteData[i]
		}
	}

	//traverse each post to get author and community info
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		//get author name from the database
		user, err := postgres.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("GetUserByID", zap.Error(err))
			continue
		}
		//get community info from the database
		community, err := postgres.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("GetCommunityDetail", zap.Error(err))
			continue
		}
		//merge the data
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteMap[post.ID],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
