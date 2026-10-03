package redis

import (
	"context"
	"lelForum/models"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func getIDsFromKey(ctx context.Context, key string, pageNum, pageSize int64) ([]string, error) {
	start := (pageNum - 1) * pageSize
	end := start + pageSize - 1
	return client.ZRevRange(ctx, key, start, end).Result()
}

func GetPostIDsInOrder(ctx context.Context, p *models.ParamPostList) ([]string, error) {
	//get post id list from redis based on the order criteria
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == "score" {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//calculate start and end index
	return getIDsFromKey(ctx, key, p.PageNum, p.PageSize)
}

func GetPostVoteData(ctx context.Context, ids []string) (data []int64, err error) {
	pipeline := client.Pipeline() //Complete all commands in the pipeline at once
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		//count the number of upvotes (score of 1) for each post
		pipeline.ZCount(ctx, key, "1", "1").Val()
	}
	cmds, err := pipeline.Exec(ctx)
	if err != nil {
		return
	}
	data = make([]int64, len(ids))
	for idx, cmd := range cmds {
		v := cmd.(*redis.IntCmd).Val()
		data[idx] = v
	}
	return
}

func GetCommunityPostIDsInOrder(ctx context.Context, p *models.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == "score" {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	// Use Redis ZINTERSTORE to get the intersection of community post IDs and score-ordered post IDs
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))
	// Use cache key to reduce zinterstore frequency
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(ctx, key).Val() < 1 {
		// Cache miss, need to create a new sorted set zinterstore
		pipeline := client.Pipeline()
		pipeline.ZInterStore(ctx, key, &redis.ZStore{
			Keys:      []string{cKey, orderKey},
			Aggregate: "MAX",
		})
		// Set expiration time for the cache key
		pipeline.Expire(ctx, key, 60*time.Second)
		_, err := pipeline.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	return getIDsFromKey(ctx, key, p.PageNum, p.PageSize)
}
