package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	Ctx       = context.Background()
	DBConfs   *simplejson.Json
	RedisPool = RedisConnector{Cli: make(map[string]*redis.Client)}
)

type RedisConnector struct {
	Cli  map[string]*redis.Client
	lock sync.RWMutex
}

func (r *RedisConnector) GetKey(key string) (*redis.Client, bool) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	con, ok := r.Cli[key]
	return con, ok
}

func (r *RedisConnector) SetKey(key string, v *redis.Client) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.Cli[key] = v
}

func (r *RedisConnector) GetCli(env, name string) (*redis.Client, error) {
	key := fmt.Sprintf("%s_%s", env, name)
	if cli, ok := r.GetKey(key); ok {
		if _, err := cli.Ping(Ctx).Result(); err == nil {
			return cli, nil
		}
	}
	cl, err := newRedisClient(env, name)
	if err != nil {
		return nil, fmt.Errorf("获取%s环境 %s redis实例失败。Error:%v", env, name, err.Error())
	}

	r.SetKey(key, cl)
	return cl, nil
}

func newRedisClient(env, redisName string) (*redis.Client, error) {
	//var (
	//once sync.Once
	//err  error
	//cli *redis.Client
	//)

	cfg, err := parseConfig(env, redisName)
	if err != nil {
		return nil, err
	}

	cli := redis.NewClient(cfg)

	_, err = cli.Ping(Ctx).Result()
	//cli.Ping(Ctx).Result()
	//if _, err := cli.Ping(Ctx).Result(); err == nil {
	//	return cli, nil
	//}
	ret, err1 := cli.Do(Ctx, `SET`, "testDo", "valueDo").Result()
	if err1 != nil {
		fmt.Printf("Do -- error:%v \n", err1.Error())
	} else {
		fmt.Printf("Do -- ret:%v \n", ret)
	}
	//for k, v := range RedisPool.Cli {
	//	fmt.Printf("RedisPool -- key:%v value:%v \n",k,v)
	//}
	return cli, err
}

func parseConfig(env, redisName string) (*redis.Options, error) {
	//fmt.Printf("getPath -- env:%s name:%s \n", env,redisName)
	//fmt.Printf("DBConfs:%v",DBConfs)
	result, err := DBConfs.GetPath(env, redisName).Map()
	if err != nil {
		fmt.Printf("getPath -- err:%s ", err.Error())
		return nil, err
	}
	server := result["Server"].(string)
	port := result["Port"].(string)
	server = fmt.Sprintf("%s:%s", server, port)
	DB, err := result["Db"].(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	pwd := result["Password"].(string)
	tm, err := result["IdleTimeout"].(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	return &redis.Options{
		Addr: server, DB: int(DB), Password: pwd, IdleTimeout: time.Duration(tm) * time.Second,
	}, nil
}

func (r *RedisConnector) Close() {
	for _, v := range r.Cli {
		v.Close()
	}
}
