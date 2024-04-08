package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"practice/settings"
)

var Client *redis.Client
var ctx = context.Background()

func Init(cfg *settings.RedisConfig) (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,     //连接池大小
		MinIdleConns: cfg.MimIdleConns, //最小空闲连接数
	})
	_, err = Client.Ping(ctx).Result()
	if errors.Is(err, redis.Nil) {
		zap.L().Error("键值不存在", zap.Any("err", err))
		return err
	}
	if err != nil {
		zap.L().Error("连接失败", zap.Any("err", err))
		return err
	}
	return nil

}
func Close() {
	if Client != nil {
		_ = Client.Close()
	}
}
