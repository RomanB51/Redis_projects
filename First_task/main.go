package main

import (
	logger "REDIS_PROJECT/log"
	"REDIS_PROJECT/redis"
	"fmt"
	"time"
)

func main() {

	logger.Init_log()
	defer logger.GetLogFile().Close()

	redis.InitClientRedis()
	redis.GetClient().SetEx(*redis.GetContext(), "test_key", "test_value", time.Second*60)
	fmt.Println(redis.GetClient().Get(*redis.GetContext(), "test_key"))
	time.Sleep(time.Second * 6)
	fmt.Println(redis.GetClient().TTL(*redis.GetContext(), "test_key"))
	fmt.Println(redis.GetClient().Exists(*redis.GetContext(), "test_key"))
	redis.GetClient().Del(*redis.GetContext(), "test_key")
	fmt.Println(redis.GetClient().Exists(*redis.GetContext(), "test_key"))

}
