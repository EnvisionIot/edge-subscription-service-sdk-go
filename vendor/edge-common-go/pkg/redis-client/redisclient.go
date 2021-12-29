package redis_client

import (
	"edge-common-go/pkg/log"
	"github.com/garyburd/redigo/redis"
	"time"
)

/**
 * @Author: qianjialin
 * @Date: 2020/8/11 14:55
 */

var pool *redis.Pool
var Log *log.LogAa

type InitParams struct {
	Network         string
	UseAuth         bool
	AuthPassword    string
	MaxIdle         int
	MaxActive       int
	IdleTimeout     time.Duration
	Wait            bool
	MaxConnLifetime time.Duration
	Host            string
	Logs            []*log.LogAa
}

func InitRedis(initParams InitParams) {
	//初始化日志
	getLog()
	// 初始化 redisclient 连接池，全局唯一
	initRedisPool(initParams)
	// 实例化全局唯一 CurrentRedisDao 用于操作RedisDao信息
	CurrentRedisDao = InitRedisDao(pool)
	var conn redis.Conn
	conn = pool.Get()
	if conn.Err() != nil {
		//重连直到成功
		initRedisPoolUntilSuccess(initParams)
	} else {
		defer func() {
			_ = conn.Close()
		}()
	}
}

func initRedisPoolUntilSuccess(params InitParams) {
	var ticker *time.Ticker
	ticker = time.NewTicker(time.Second * 3)

	var err error

	for {
		select {
		case <-ticker.C:
			var conn redis.Conn
			conn = pool.Get()
			if conn.Err() == nil {
				defer func() {
					_ = conn.Close()
				}()
				ticker.Stop()
				return
			} else {
				Log.Errorf("initRedisPool error, err=%s", err)
			}
		}
	}
}

func initRedisPool(params InitParams) {
	pool = &redis.Pool{
		// 初始化链接数量
		MaxIdle:         params.MaxIdle,
		MaxActive:       params.MaxActive,
		IdleTimeout:     params.IdleTimeout,
		Wait:            params.Wait,
		MaxConnLifetime: params.MaxConnLifetime,
		Dial: func() (redis.Conn, error) {
			var c redis.Conn
			var err error
			c, err = redis.Dial(params.Network, params.Host)
			if err != nil {
				return nil, err
			}
			if params.UseAuth {
				_, err = c.Do("AUTH", params.AuthPassword)
				if err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
}

func getLog() {
	Log, _ = log.GetLogAa("/redisclient/redisclient")
}
