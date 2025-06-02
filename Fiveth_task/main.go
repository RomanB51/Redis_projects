package main

import (
	logger "fiveth_task/log"
	"fiveth_task/redis"
	"fmt"
)

func main() {
	logger.Init_log()
	redis.InitClientRedis()
	fmt.Println(redis.Rdb.SAdd(redis.Ctx, "article:123:tags", "python", "redis", "db", "nosql", "cache"))
	fmt.Println(redis.Rdb.SIsMember(redis.Ctx, "article:123:tags", "python"))
	fmt.Println(redis.Rdb.SAdd(redis.Ctx, "popular_tags", "nosql"))
	fmt.Println(redis.Rdb.SInter(redis.Ctx, "article:123:tags", "popular_tags"))
	fmt.Println(redis.Rdb.SRem(redis.Ctx, "article:123:tags", "python", "redis"))
	fmt.Println(redis.Rdb.SPop(redis.Ctx, "article:123:tags"))
}
