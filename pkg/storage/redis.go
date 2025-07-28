package storage

import "github.com/redis/go-redis/v9"

func NewRedis(dsn string) (*redis.Client, error) {
	opts, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	return redis.NewClient(opts), nil
}
