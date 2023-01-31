package dao

import (
	"context"
	"cool/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisParameter struct {
	Env  string        `json:"env" binding:"required"`
	Name string        `json:"name" binding:"required"`
	Args []interface{} `json:"args" binding:"required"`
}

var ctx = context.Background()

type IRedisManger interface {
	Set(key string, value interface{}) (string, error)
}
type RedisManager struct {
	Env  string
	Name string
}

func NewRedisManger(env, name string) IRedisManger {
	return &RedisManager{env, name}
}

func getRedis(env, name string) (*redis.Client, error) {
	cl, err := config.RedisPool.GetCli(env, name)
	if err != nil {
		return nil, fmt.Errorf("在redis-dao中获取%s环境 %s redis实例失败。Error:%v", env, name, err.Error())
	}

	return cl, err
}

// Set key value
func (c *RedisManager) Set(key string, value interface{}) (string, error) {
	cli, err := getRedis(c.Env, c.Name)
	if err != nil {
		return "", err
	}
	ret := cli.Do(ctx, `SET`, key, value)
	if ret.Err() != nil {
		log.Println(ret.Err())
		return "", ret.Err()
	}
	return ret.String(), nil
}

// Get key value
func Get(env, name, key string) (string, error) {
	cli, err := getRedis(env, name)
	if err != nil {
		return "", err
	}
	ret := cli.Do(ctx, `GET`, key)
	if ret.Err() != nil {
		log.Println(ret.Err())
		return "", ret.Err()
	}
	return ret.String(), nil
}

//Hget key field
func Hget(env, name, key, field string) (string, error) {
	cli, err := getRedis(env, name)
	if err != nil {
		return "", err
	}
	ret := cli.Do(ctx, `HGET`, key, field)
	if ret.Err() != nil {
		log.Println(ret.Err())
		return "", ret.Err()
	}
	return ret.String(), nil
}

//Hset key field value
func Hset(env, name, key, field string, value interface{}) (string, error) {
	cli, err := getRedis(env, name)
	if err != nil {
		return "", err
	}
	ret := cli.Do(ctx, `HSET`, key, field, value)
	if ret.Err() != nil {
		log.Println(ret.Err())
		return "", ret.Err()
	}
	return ret.String(), nil
}
