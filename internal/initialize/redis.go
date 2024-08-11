package initialize

import (
	"context"
	"fmt"

	"github.com/hiumx/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB,
		PoolSize: r.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Initialize Redis failed!", zap.Error(err))
		return
	}
	global.Logger.Info("Initialize Redis successfully!")
	global.Rdb = rdb

	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "test", 200, 0).Err()
	if err != nil {
		fmt.Println("Error set example test Redis: ", err)
		return
	}

	val, err := global.Rdb.Get(ctx, "test").Result()
	if err != nil {
		fmt.Println("Error get example test Redis: ", err)
		return
	}

	fmt.Println("Result get of test example Redis:::", val)
}
