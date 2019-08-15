package tasks

import "github.com/tiancai110a/go-play/goplay-bk/cache"

// RestartDailyRank 重启一天的排名
func RestartDailyRank() error {
	return cache.RedisClient.Del("rank:daily").Err()
}
