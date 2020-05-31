package short

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
	"math/rand"
	"strconv"
	"sync"
)

type RedisDB struct {
	ctx    context.Context
	client *redis.Client
}

// new redisdb
func NewRedisDB(options *redis.Options) *RedisDB {
	ctx := context.Background()
	redisdb := new(RedisDB)
	client := redis.NewClient(options)
	redisdb.client = client
	redisdb.ctx = ctx
	once := new(sync.Once)
	once.Do(func() {
		pong, err := client.Ping(ctx).Result()
		fmt.Println(pong, err)
	})
	return redisdb
}

func (rd *RedisDB) Add(v *Node) (shortURL string, err error) {
	if _,err := rd.client.Get(rd.ctx,v.ShortValue).Result();err !=redis.Nil {
		v.ShortValue = v.ShortValue+strconv.FormatInt(int64(rand.Int()),10)
	}
	//
	if err := rd.client.Set(rd.ctx, v.ShortValue, v.LongValue, 0).Err(); err != nil {
		return "", err
	} else {
		return v.ShortValue, nil
	}
}

func (rd *RedisDB) Delete(shortURL string) {
	if _, err := rd.client.Del(rd.ctx, shortURL).Result(); err != nil {
		glog.Error(err)
	}
}
func (rd *RedisDB) Change(vi *Node, r string) error {
	long, err := rd.client.Get(rd.ctx, vi.ShortValue).Result()
	if err != nil {
		return err
	}
	if err := rd.client.Set(rd.ctx, r, long, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (rd *RedisDB) Find(shortURL string) (string, error) {
	long, err := rd.client.Get(rd.ctx, shortURL).Result()
	if err != nil {
		return "", err
	}
	return long, nil
}
