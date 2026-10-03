package redis

import (
	"context"
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	OneWeekInSeconds = float64(7 * 24 * 3600)
)

var (
	ErrVoteTimeExpire = errors.New("vote time expired")
	VoteScore         = float64(432) // One upvote adds 432 points to the post's score
	ErrVoteRepeated   = errors.New("repeated vote")
)

func CreatePost(ctx context.Context, postID uint64, communityID int64) (err error) {
	//Use pipeline to execute multiple commands as a single transaction
	pipeline := client.TxPipeline()
	//Post time
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//Initialize post score
	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  0,
		Member: postID,
	})
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(ctx, cKey, postID)
	_, err = pipeline.Exec(ctx)
	if err != nil {
		zap.L().Error("redis post failed", zap.Error(err))
	}
	return
}

func VoteForPost(ctx context.Context, userID, postID string, direction float64) (err error) {
	//Judge the voting situation
	//Retrieve the post's creation time
	postTime := client.ZScore(ctx, getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds {
		zap.L().Debug("VoteForPost time expired", zap.String("postID", postID))
		return ErrVoteTimeExpire
	}
	//Make change to vote
	//Search for the previous vote record
	originalVoteVal := client.ZScore(ctx, getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	//If the vote direction is the same as before, no need to process
	if direction == originalVoteVal {
		return ErrVoteRepeated
	}
	var dir float64
	if direction > originalVoteVal {
		dir = 1
	} else {
		dir = -1
	}
	//Calculate the difference of votes
	diff := math.Abs(direction - originalVoteVal)
	//The following operations need to be atomic(in a transaction)
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), dir*diff*VoteScore, postID)

	//Record user's voting behavior
	if direction == 0 {
		//Remove the record
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  direction,
			Member: userID,
		})
	}
	_, err = pipeline.Exec(ctx)
	if err != nil {
		zap.L().Debug("VoteForPost vote redis failed", zap.String("postID", postID))
	}
	return
}
