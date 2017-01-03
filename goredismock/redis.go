package goredismock

import (
	"github.com/stretchr/testify/mock"
)

type RedisTemplate struct {
	mock.Mock
}

func (r RedisTemplate) Get(key string) ([]byte, error) {
	argsMock := r.Called(key)

	return []byte(argsMock.String(0)), argsMock.Error(1)
}

func (r RedisTemplate) Setex(key string, seconds int, value string) error {
	argsMock := r.Called(key, seconds, value)

	return argsMock.Error(0)
}

func (r RedisTemplate) SimpleSet(key, value string) error {
	argsMock := r.Called(key, value)

	return argsMock.Error(0)
}
