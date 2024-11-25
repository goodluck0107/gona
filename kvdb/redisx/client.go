package redisx

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"

	"github.com/go-redis/redis"
)

type (
	// ClientType is the type of client instance
	clientPool struct {
		options *redis.Options
		client  *redis.Client
	}
)

var clientPools = make(map[string]*clientPool)

// GetClient is to get a client from redis pool
// Funcs of redis.Client are thread-safe
func GetClient(options *redis.Options) (*redis.Client, error) {
	pool, err := getClientPool(options)
	if err != nil {
		return nil, err
	}
	return pool.client, nil
}

func getClientPool(options *redis.Options) (*clientPool, error) {
	typ := getType(options)
	pool, ok := clientPools[typ]
	if ok {
		return pool, nil
	}

	InitClientPool(options)

	pool, ok = clientPools[typ]
	if !ok {
		return nil, fmt.Errorf("cannot find redis client type %s", typ)
	}
	return pool, nil
}

// InitClientPool Init one redis pool
// It will do nothing if it's repeated init.
func InitClientPool(options *redis.Options) {
	typ := getType(options)
	if _, ok := clientPools[typ]; ok {
		return
	}
	logger.Debug("redis", options.Addr, options.DB, options.PoolSize, options.MinIdleConns, options.MaxRetries)
	client := redis.NewClient(options)
	clientPools[typ] = &clientPool{
		options: options,
		client:  client,
	}
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
}

// NewClientOptions create new client options by client type
func getType(options *redis.Options) string {
	return fmt.Sprintf("%v_%v", options.Addr, options.DB)
}
