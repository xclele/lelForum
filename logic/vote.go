package logic

import (
	"lelForum/database/redis"
	"lelForum/models"
	"strconv"

	"go.uber.org/zap"
)

// Simple implementation of voting logic
// Adding 432 points to the post's score for an upvote (1 day)
/*
Cases:
Dir = 1: User upvotes
	1. User votes for the first time	Diff: 1
	2. User changes from downvote to upvote	Diff: 2
Dir = 0: User cancels vote
	1. User cancels upvote
	2. User cancels downvote
Dir = -1: User downvotes
	1. User downvotes for the first time
	2. User changes from upvote to downvote
Cons:
A post can only be voted on within 7 days of its creation
	1. After 7 days, Save the vote into the database
	2. Remove the vote data from KeyPostVotedZSetPF
*/
func VoteForPost(userID uint64, p *models.ParamVoteData) (err error) {
	zap.L().Debug("VoteForPost",
		zap.Uint64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.FormatUint(userID, 10), p.PostID, float64(p.Direction))
}
