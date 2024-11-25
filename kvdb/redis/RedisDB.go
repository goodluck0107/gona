package redis

type RedisDB struct {
	Client *Client
}

var RedisTimeDefaultValue string = "" //value string 空值
func NewRedisDB(addr string, db int, pwd string) (redisDB *RedisDB, err error) {
	redisDB = &RedisDB{}
	redisDB.Client = &Client{
		Addr:        addr,
		Db:          db, // default db is 0
		Password:    pwd,
		MaxPoolSize: 5,
	}
	if err = redisDB.Client.Auth(pwd); err != nil {
		return
	}
	return
}
