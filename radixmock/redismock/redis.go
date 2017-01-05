package redismock

import (
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/stretchr/testify/mock"
)

type RedisTemplate struct {
	mock.Mock
}

func (r RedisTemplate) Cmd(cmd string, args ...interface{}) *redis.Resp {
	argsMock := r.Called(cmd, args)

	return argsMock.Get(0).(*redis.Resp)
}
