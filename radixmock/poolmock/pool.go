package poolmock

import (
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/stretchr/testify/mock"
)

type PoolTemplate struct {
	mock.Mock
}

func (p PoolTemplate) Cmds(cmd string, args ...interface{}) *redis.Resp {
	argsMock := p.Called(cmd, args)

	return argsMock.Get(0).(*redis.Resp)
}
