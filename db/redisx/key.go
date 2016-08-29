package redisx

import (
	"github.com/garyburd/redigo/redis"
)

func (p *Redis) Del(k ...interface{}) (int, error) {
	return redis.Int(p.Do("DEL", k...))
}
func Del(k ...interface{}) (int, error) { return Default.Del(k...) }
