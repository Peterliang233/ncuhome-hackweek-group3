package dao

import "github.com/garyburd/redigo/redis"

var (
	Conn redis.Conn
)