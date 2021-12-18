// Package cache assemble all functions for redis
package cache

import (
	"fmt"
	"regexp"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

// RedisSet permit to put k/v in redis
func RedisSet(redisURI string, keyName string, value string, expire time.Duration) {
	var (
		ctx = context.Background()
	)
	opt, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Error().Err(err).Msgf("Error occured while connecting to redis address %s", opt.Addr)
		return
	}

	rdb := redis.NewClient(opt)
	err = rdb.SetNX(ctx, keyName, value, expire*time.Second).Err()
	if err != nil {
		log.Error().Err(err).Msgf("Failed to set keyName on redis address %s", opt.Addr)
	}
}

// RedisGet permit to get k/v in redis
func RedisGet(redisURI string, keyName string) (z string, err error) {
	var (
		ctx = context.Background()
		opt *redis.Options
	)
	opt, err = redis.ParseURL(redisURI)
	if err != nil {
		log.Error().Err(err).Msgf("Error occured while connecting to redis address %s", opt.Addr)
		return
	}

	rdb := redis.NewClient(opt)
	z, err = rdb.Get(ctx, keyName).Result()
	return
}

// RedisGet permit to get k/v in redis
func RedisDel(redisURI string, keyName string) {
	var (
		ctx = context.Background()
	)
	opt, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Error().Err(err).Msgf("Error occured while connecting to redis address %s", opt.Addr)
		return
	}

	rdb := redis.NewClient(opt)
	err = rdb.Del(ctx, keyName).Err()
	if err != nil {
		log.Warn().Err(err).Msgf("Error occured while deleting keyName %s on redis address %s", keyName, opt.Addr)
	}
}

// RedisGet permit to get k/v in redis
func RedisKeys(redisURI string, keyPrefix string) (z []string) {
	var (
		ctx = context.Background()
	)
	opt, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Error().Err(err).Msgf("Error occured while connecting to redis address %s", opt.Addr)
		return
	}

	rdb := redis.NewClient(opt)
	z, err = rdb.Keys(ctx, keyPrefix).Result()
	if err != nil {
		log.Warn().Err(err).Msgf("Error occured while finding keys with prefix %s on redis address %s", keyPrefix, opt.Addr)
	}
	return
}

// RedisDelWithPrefix permit to list keys keys has prefix like keyName* and then, delete keys
func RedisDelWithPrefix(redisURI string, keyPrefix string) {
	var (
		ctx = context.Background()
	)
	opt, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Error().Err(err).Msgf("Error occured while connecting to redis address %s", opt.Addr)
		return
	}
	matched, _ := regexp.MatchString(`\*$`, keyPrefix)
	if !matched {
		log.Debug().Msgf("keyPrefix %s does not finish with wildcard, so let's set it", keyPrefix)
		keyPrefix = fmt.Sprintf("%s*", keyPrefix)
	}
	keys := RedisKeys(redisURI, keyPrefix)
	if len(keys) > 0 {
		rdb := redis.NewClient(opt)
		for _, keyName := range keys {
			err = rdb.Del(ctx, keyName).Err()
			if err != nil {
				log.Warn().Err(err).Msgf("Error occured while deleting keyName %s on redis address %s", keyName, opt.Addr)
			}
		}
	}
}

// RedisDeleteKeysHasPrefix permit to list keys keys has prefix like keyName* and then, delete keys
func RedisDeleteKeysHasPrefix(redisURI string, prefixes []string) {
	if len(prefixes) > 0 {
		for _, keyPrefix := range prefixes {
			var reg = regexp.MustCompile(`\*$`)

			if !reg.MatchString(keyPrefix) {
				log.Debug().Msgf("keyPrefix %s does not finish with wildcard, so let's set it", keyPrefix)
				keyPrefix = fmt.Sprintf("%s*", keyPrefix)
			}
			keys := RedisKeys(redisURI, keyPrefix)
			if len(keys) > 0 {
				for _, key := range keys {
					RedisDel(redisURI, key)
				}
			}
		}
	}
}

// RedisFlushDB permit to flush actual used db in redis
func RedisFlushDB(redisURI string) {
	var (
		ctx = context.Background()
	)
	opt, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Error().Err(err).Msgf("Error occured while connecting to redis address %s", opt.Addr)
		return
	}

	rdb := redis.NewClient(opt)
	err = rdb.FlushDB(ctx).Err()
	if err != nil {
		log.Warn().Err(err).Msgf("Error occured while flushing DB on redis address %s", opt.Addr)
	}
}

// RedisFlushAll permit to flush all db in redis
func RedisFlushAll(redisURI string) {
	var (
		ctx = context.Background()
	)
	opt, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Error().Err(err).Msgf("Error occured while connecting to redis address %s", opt.Addr)
		return
	}

	rdb := redis.NewClient(opt)
	err = rdb.FlushAll(ctx).Err()
	if err != nil {
		log.Warn().Err(err).Msgf("Error occured while flushing all DBs on redis address %s", opt.Addr)
	}
}

// RedisPing permit to get redis status
func RedisPing(redisURI string) (b bool) {
	var (
		ctx = context.Background()
	)
	opt, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Error().Err(err).Msgf("Error occured while connecting to redis address %s", opt.Addr)
		return
	}

	rdb := redis.NewClient(opt)
	err = rdb.Ping(ctx).Err()
	if err != nil {
		log.Error().Err(err).Msgf("Error occured while pinging redis %s address", opt.Addr)
		return
	}
	return true
}
