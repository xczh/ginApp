package redis

import (
	"app/internal/config"
	"app/internal/log"
	"context"
	"errors"
	"runtime"
	"time"

	goredis "github.com/go-redis/redis/v8"
)

var _redisClient *goredis.Client

func Initialize() error {
	rdb, err := newRedisClient()
	if rdb == nil {
		return err
	}
	_redisClient = rdb
	return nil
}

func Terminate() {
	if _redisClient == nil {
		return
	}
	if err := _redisClient.Close(); err != nil {
		log.Logger().Error("Close Redis client error: ", err)
		return
	}
	_redisClient = nil
}

func newRedisClient() (*goredis.Client, error) {
	dsn := config.GetRedisDSN()
	if dsn == "" {
		return nil, errors.New("Empty Redis DSN")
	}
	opts, err := goredis.ParseURL(dsn)
	if err != nil {
		log.Logger().Error("Invalid Redis DSN: ", err)
		return nil, err
	}
	opts.MinIdleConns = runtime.NumCPU()
	rdb := goredis.NewClient(opts)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Logger().Error("Connect to Redis server error: ", err)
		return nil, err
	}
	return rdb, nil
}
