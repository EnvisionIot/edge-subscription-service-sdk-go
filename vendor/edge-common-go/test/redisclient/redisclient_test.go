package redisclient

import (
	"edge-common-go/pkg/log"
	"edge-common-go/pkg/redis-client"
	"edge-common-go/pkg/utils"
	"testing"
	"time"
)

/**
 * @Author: qianjialin
 * @Date: 2020/12/25 9:50
 */
func Test_redisClientInit(t *testing.T) {
	log.LogAGlobal.Init(utils.GetBaseDir())
	var initParams = redis_client.InitParams{}
	initParams.Network = "tcp"
	initParams.UseAuth = false
	initParams.MaxIdle = 16
	initParams.MaxActive = 0
	initParams.IdleTimeout = time.Duration(300000) * time.Millisecond
	initParams.Wait = true
	initParams.MaxConnLifetime = time.Duration(0) * time.Millisecond
	initParams.Host = "127.0.0.1:6379"
	//redis_client.InitRedis(initParams)
}
