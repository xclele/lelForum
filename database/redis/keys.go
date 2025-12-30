package redis

// Redis keys
const (
	KeyPrefix          = "lelforum:"   // Key prefix
	KeyPostTimeZSet    = "post:time"   // Sorted set for post times
	KeyPostScoreZSet   = "post:score"  // Sorted set for post scores (Total votes)
	KeyPostVotedZSetPF = "post:voted:" // Prefix for sorted sets tracking users who voted on a post
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
