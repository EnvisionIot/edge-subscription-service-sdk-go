package utils

import (
	"edge-common-go/pkg/httpclient"
	"edge-common-go/pkg/log"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

/**
 * Description:
 *
 * @author yang.zhang4
 * @date 2020/12/23
 */

var (
	fileName = GetFileNameLastN(1)
)

//连接消费数据的nsqd，失败时重试指定次数，retryIntervalMs是两次重连之间的时间间隔，retryTimes是重连次数
func ConnectNSQDSeveralTimes(logs []*log.LogAa, consumer *nsq.Consumer, addr string, retryIntervalMs int64, retryTimes int64) error {
	var err error
	err = consumer.ConnectToNSQD(addr)
	if err == nil {
		return nil
	}

	log.MultiLogErrorf(logs, "[%s:%d]:connectToNSQD(%s), %s", fileName, GetLineNum(), addr, err)

	if retryTimes <= 0 {
		return err
	}

	var ticker *time.Ticker
	ticker = time.NewTicker(time.Millisecond * time.Duration(retryIntervalMs))

	var count int64 = 0
	for {
		select {
		case <-ticker.C:
			err = consumer.ConnectToNSQD(addr)
			if err == nil {
				ticker.Stop()
				return nil
			}
			log.MultiLogErrorf(logs, "[%s:%d]:connectToNSQD(%s), %s", fileName, GetLineNum(), addr, err)
			count++
			if count >= retryTimes {
				ticker.Stop()
				return err
			}
		}
	}
}

//连接消费数据的nsqd直到成功，retryIntervalMs是两次重连之间的时间间隔
func ConnectNSQDUntilSuccess(logs []*log.LogAa, consumer *nsq.Consumer, addr string, retryIntervalMs int64) error {
	var err error
	err = consumer.ConnectToNSQD(addr)
	if err == nil {
		return nil
	}

	var ticker *time.Ticker
	ticker = time.NewTicker(time.Millisecond * time.Duration(retryIntervalMs))

	for {
		select {
		case <-ticker.C:
			err = consumer.ConnectToNSQD(addr)
			if err == nil {
				ticker.Stop()
				return nil
			}
			log.MultiLogErrorf(logs, "[%s:%d]:connectToNSQD(%s), %s", fileName, GetLineNum(), addr, err)
		}
	}
}

//创建topic
func CreateNSQTopic(url string, topicName string, webOpTimeOutMs int64) error {
	var err error
	var httpUrl = fmt.Sprintf("%s?topic=%s", url, topicName)
	_, err = httpclient.HttpDefaultGetWithTimeout(httpUrl, webOpTimeOutMs)
	if err != nil {
		return err
	}
	return nil
}

//创建topic，失败时重试指定次数，retryIntervalMs是两次重试之间的时间间隔，retryTimes是重试次数
func CreateNSQTopicSeveralTimes(logs []*log.LogAa, url string, topicName string, webOpTimeOutMs int64, retryIntervalMs int64, retryTimes int64) error {
	var err error
	err = CreateNSQTopic(url, topicName, webOpTimeOutMs)
	if err == nil {
		return nil
	}

	log.MultiLogErrorf(logs, "[%s:%d]:createNSQTopic(%s, %s, %d), %s", fileName, GetLineNum(), url, topicName, webOpTimeOutMs, err)

	if retryTimes <= 0 {
		return err
	}

	var ticker *time.Ticker
	ticker = time.NewTicker(time.Millisecond * time.Duration(retryIntervalMs))

	var count int64 = 0
	for {
		select {
		case <-ticker.C:
			err = CreateNSQTopic(url, topicName, webOpTimeOutMs)
			if err == nil {
				ticker.Stop()
				return nil
			}
			log.MultiLogErrorf(logs, "[%s:%d]:createNSQTopic(%s, %s, %d), %s", fileName, GetLineNum(), url, topicName, webOpTimeOutMs, err)
			count++
			if count >= retryTimes {
				ticker.Stop()
				return err
			}
		}
	}
}

//设定topic超时时间
func SetNSQTopicTimeout(url string, topicName string, timeoutSec int64, webOpTimeOutMs int64) error {
	var err error
	//var responseByteArray []byte
	var httpUrl = fmt.Sprintf("%s?topic=%s&ts=%d&d=true", url, topicName, timeoutSec)
	_, err = httpclient.HttpDefaultGetWithTimeout(httpUrl, webOpTimeOutMs)
	if err != nil {
		return err
	}
	return nil
}

//设定topic超时时间，失败时重试指定次数，retryIntervalMs是两次重试之间的时间间隔，retryTimes是重试次数
func SetNSQTopicTimeoutSeveralTimes(logs []*log.LogAa, url string, topicName string, timeoutSec int64, webOpTimeOutMs int64, retryIntervalMs int64, retryTimes int64) error {
	var err error
	err = SetNSQTopicTimeout(url, topicName, timeoutSec, webOpTimeOutMs)
	if err == nil {
		return nil
	}

	log.MultiLogErrorf(logs, "[%s:%d]:setNSQTopicTimeout(%s, %s, %d, %d), %s", fileName, GetLineNum(), url, topicName, timeoutSec, webOpTimeOutMs, err)

	if retryTimes <= 0 {
		return err
	}

	var ticker *time.Ticker
	ticker = time.NewTicker(time.Millisecond * time.Duration(retryIntervalMs))

	var count int64 = 0
	for {
		select {
		case <-ticker.C:
			err = SetNSQTopicTimeout(url, topicName, timeoutSec, webOpTimeOutMs)
			if err == nil {
				ticker.Stop()
				return nil
			}
			log.MultiLogErrorf(logs, "[%s:%d]:setNSQTopicTimeout(%s, %s, %d, %d), %s", fileName, GetLineNum(), url, topicName, timeoutSec, webOpTimeOutMs, err)
			count++
			if count >= retryTimes {
				ticker.Stop()
				return err
			}
		}
	}
}

//ping某个ip的nsqd
func PingNSQD(url string, webOpTimeOutMs int64) error {
	var err error
	var responseByteArray []byte
	var httpUrl = url
	responseByteArray, err = httpclient.HttpDefaultGetWithTimeout(httpUrl, webOpTimeOutMs)
	if err != nil {
		return err
	}

	var responseStr = string(responseByteArray)
	if responseStr != "OK" {
		return fmt.Errorf("ping ret %s(need OK) error", responseStr)
	}
	return nil
}

//设定topic的保留文件的数量
func SetNSQTopicRetention(url string, topicName string, retentionCount int64, webOpTimeOutMs int64) error {
	var err error
	//var responseByteArray []byte
	var httpUrl = fmt.Sprintf("%s?topic=%s&count=%d", url, topicName, retentionCount)
	_, err = httpclient.HttpDefaultGetWithTimeout(httpUrl, webOpTimeOutMs)
	if err != nil {
		return err
	}
	return nil
}

//设定topic的保留文件的数量，失败时重试指定次数，retryIntervalMs是两次重试之间的时间间隔，retryTimes是重试次数
func SetNSQTopicRetentionSeveralTimes(logs []*log.LogAa, url string, topicName string, retentionCount int64, webOpTimeOutMs int64, retryIntervalMs int64, retryTimes int64) error {
	var err error
	err = SetNSQTopicRetention(url, topicName, retentionCount, webOpTimeOutMs)
	if err == nil {
		return nil
	}

	log.MultiLogErrorf(logs, "[%s:%d]:setNSQTopicRetention(%s, %s, %d, %d), %s", fileName, GetLineNum(), url, topicName, retentionCount, webOpTimeOutMs, err)

	if retryTimes <= 0 {
		return err
	}

	var ticker *time.Ticker
	ticker = time.NewTicker(time.Millisecond * time.Duration(retryIntervalMs))

	var count int64 = 0
	for {
		select {
		case <-ticker.C:
			err = SetNSQTopicRetention(url, topicName, retentionCount, webOpTimeOutMs)
			if err == nil {
				ticker.Stop()
				return nil
			}
			log.MultiLogErrorf(logs, "[%s:%d]:setNSQTopicRetention(%s, %s, %d, %d), %s", fileName, GetLineNum(), url, topicName, retentionCount, webOpTimeOutMs, err)
			count++
			if count >= retryTimes {
				ticker.Stop()
				return err
			}
		}
	}
}
