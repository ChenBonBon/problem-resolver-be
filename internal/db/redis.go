package db

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnectRedis() {
	redisPort := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PASS")
	redisDbName := os.Getenv("REDIS_DB_NAME")

	db, err := strconv.Atoi(redisDbName)

	if err != nil {
		slog.Error("Redis DB 不存在", err.Error())
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + redisPort,
		Password: redisPass,
		DB:       db,
	})

	Redis = rdb

	slog.Info("Redis 连接成功")
}
