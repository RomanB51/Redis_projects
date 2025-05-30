package redis

import (
	"encoding/json"
	"os"
	"time"

	logger "second_tack/log"

	"context"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Ctx context.Context
)

func InitClientRedis() {
	file, err := os.Open("config/config.json")
	if err != nil {
		logger.Logger.Print("Не удалось считать конфигурационный файл: ", err)
		return
	} else {
		logger.Logger.Print("Конфигурационный файл считан")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := redis.Options{}
	err = decoder.Decode(&config)
	if err != nil {
		logger.Logger.Print("Error decoding config:", err)
		return
	}

	Ctx = context.Background()

	maxRetries := 5
	for counter := 1; counter <= maxRetries; counter++ {
		Rdb = redis.NewClient(&config)
		if _, err = Rdb.Ping(Ctx).Result(); err != nil {
			if counter == maxRetries {
				logger.Logger.Fatal("Не удалось подключиться к БД: ", err, counter)
			}
		} else {
			logger.Logger.Printf("Успешное подключение к БД c %d попытки", counter)
			break
		}
		time.Sleep(time.Second * 5)
	}

}
