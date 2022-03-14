package redis

import (
	"fmt"
	"ziweiMemo/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s%s",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.DB,       // 使用了哪个redis库
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	// 测试连接
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
