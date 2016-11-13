package redisx

import (
	"log"
	"time"
)

const LOCK = "lock:"

func DelLock(key string) (err error) { return Default.DelLock(key) }
func (p *Redis) DelLock(key string) (err error) {
	_, err = p.Del(LOCK + key)
	return
}

func GetLock(key string, timeout int64, tryCount int) (ok bool) {
	return Default.GetLock(key, timeout, tryCount)
}
func (p *Redis) GetLock(key string, timeout int64, tryCount int) (ok bool) {
	now := time.Now().Unix()
	val := now + timeout + 1
	ok, _ = p.Setnx(LOCK+key, val)
	log.Println(ok)
	if ok {
		return
	}
	for i := 0; i < tryCount; i++ {
		if ok = p.tryGetLock(key, timeout); ok {
			return true
		}
		time.Sleep(time.Second)
	}
	return p.tryGetLock(key, timeout)
}

func (p *Redis) tryGetLock(key string, timeout int64) (ok bool) {
	val, err := Int64(p.Get(LOCK + key))
	log.Println(val, err)

	if val >= time.Now().Unix() {
		return false
	}

	val2, err := Int64(p.GetSet(LOCK+key, time.Now().Unix()+timeout+1))
	log.Println(val2, err)
	return val == val2
}
