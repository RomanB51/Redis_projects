package main

import (
	"fmt"
	logger "fourth_task/log"
	"fourth_task/redis"
)

func main() {
	logger.Init_log()
	redis.InitClientRedis()
	redis.Rdb.LPush(redis.Ctx, "tasks:queue", "task1", "task2", "task3")
	fmt.Println(redis.Rdb.LRange(redis.Ctx, "tasks:queue", 1, 2))
	fmt.Println(redis.Rdb.RPop(redis.Ctx, "tasks:queue"))
	fmt.Println(redis.Rdb.LTrim(redis.Ctx, "tasks:queue", 1, 2))
	fmt.Println(redis.Rdb.RPop(redis.Ctx, "tasks:queue"))

}
