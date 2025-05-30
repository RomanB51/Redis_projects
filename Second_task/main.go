package main

import (
	"fmt"
	logger "second_tack/log"
	"second_tack/redis"
	"time"
)

func main() {
	logger.Init_log()
	redis.InitClientRedis()
	redis.Rdb.SetEx(redis.Ctx, "weather:moscow", "+25°C", time.Hour)
	time.Sleep(time.Second * 8)
	fmt.Println(redis.Rdb.TTL(redis.Ctx, "weather:moscow"))
	time.Sleep(time.Second * 2)
	redis.Rdb.SetEx(redis.Ctx, "weather:moscow", "+28°C", 0)
	fmt.Println(redis.Rdb.TTL(redis.Ctx, "weather:moscow"))
	fmt.Println(redis.Rdb.Get(redis.Ctx, "weather:moscow"))
	redis.Rdb.Del(redis.Ctx, "weather:moscow")

}
