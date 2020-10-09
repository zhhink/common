package redis

import (
	"github.com/go-redis/redis/v8"
)

// SimpleClient
func SimpleClient(addr string, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return rdb
}

// todo
// func HSetStruct(ctx context.Context, key string, values ...interface{}) *IntCmd {
// 	args := make([]interface{}, 2, 2+len(values))
// 	args[0] = "hset"
// 	args[1] = key
// 	args = appendArgs(args, values)
// 	cmd := NewIntCmd(ctx, args...)
// 	_ = c(ctx, cmd)
// 	return cmd
// }

// func HGetStruct()
