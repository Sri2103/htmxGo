package userModel

import (
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
)

type Session struct {
}

func LoadSession() *scs.SessionManager {
	// TODO: Impl

	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "host:6379")
		},
	}
	sessionManager := scs.New()

	sessionManager.Store = redisstore.New(pool)

	return sessionManager

}
