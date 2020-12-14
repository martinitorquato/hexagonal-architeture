package mock_redis

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

// Connection mocks redis for unit test
func Connection() (*redis.Client, *miniredis.Miniredis) {
	miniredis, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: miniredis.Addr(),
		DB:   0,
	})
	return client, miniredis
}
