package redis

// Redis keys
const (
	KeyPrefix          = "lelforum:"   // Key prefix
	KeyPostTimeZSet    = "post:time"   // Sorted set for post times
	KeyPostScoreZSet   = "post:score"  // Sorted set for post scores (Total votes)
	KeyPostVotedZSetPF = "post:voted:" // Prefix for sorted sets tracking users who voted on a post
	KeyCommunitySetPF  = "community:"  // Prefix for sets of post IDs in a community
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
