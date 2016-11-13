package redisx

import (
	"github.com/eynstudio/gobreak"
	"github.com/garyburd/redigo/redis"
)

func (p *Redis) Set(k string, v gobreak.T) (string, error) { return redis.String(p.Do("SET", k, v)) }
func Set(k string, v gobreak.T) (string, error)            { return Default.Set(k, v) }

func (p *Redis) Setnx(k string, v gobreak.T) (bool, error) { return redis.Bool(p.Do("SETNX", k, v)) }
func Setnx(k string, v gobreak.T) (bool, error)            { return Default.Setnx(k, v) }

func (p *Redis) GetSet(k string, v gobreak.T) (interface{}, error) { return p.Do("GETSET", k, v) }
func GetSet(k string, v gobreak.T) (interface{}, error)            { return Default.GetSet(k, v) }

func (p *Redis) Get(k string) (interface{}, error) { return p.Do("GET", k) }
func Get(k string) (interface{}, error)            { return Default.Get(k) }
func (p *Redis) GetStr(k string) (string, error)   { return redis.String(p.Get(k)) }
func GetStr(k string) (string, error)              { return Default.GetStr(k) }
