package redis

import (
	"encoding/json"
	"os"

	logger "REDIS_PROJECT/log"

	"github.com/redis/go-redis/v9"

	"context"
)

var (
	rdb *redis.Client
	ctx context.Context
)

func InitClientRedis() {

	file, err := os.Open("config/config.json")
	if err != nil {
		logger.PrintErr("Не удалось считать конфигурационный файл: ", err)
		return
	} else {
		logger.Print("Конфигурационный файл считан")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := redis.Options{}
	err = decoder.Decode(&config)
	if err != nil {
		logger.Log.Println("Error decoding config:", err)
		return
	}

	ctx = context.Background()

	rdb = redis.NewClient(&config)

	if _, err = rdb.Ping(ctx).Result(); err != nil {
		logger.PrintErr("Ошибка подключения к БД: ", err)
	} else {
		logger.Print("Успешное подключение к БД")
	}

}

func GetClient() *redis.Client {
	return rdb
}

func GetContext() *context.Context {
	return &ctx
}
