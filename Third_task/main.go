package main

import (
	logger "Third_task/log"
	"Third_task/redis"
	"fmt"
)

func main() {
	logger.Init_log()
	redis.InitClientRedis()
	redis.Rdb.HSet(redis.Ctx, "user:1000", "name", "John Doe", "email", "john@example.com", "login_count", 0)
	redis.Rdb.HIncrBy(redis.Ctx, "user:1000", "login_count", 1)
	fmt.Println(redis.Rdb.HGetAll(redis.Ctx, "user:1000"))
	redis.Rdb.HDel(redis.Ctx, "user:1000", "name")
	fmt.Println(redis.Rdb.HGetAll(redis.Ctx, "user:1000"))

}
