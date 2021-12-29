package nsq_client

import (
	"edge-common-go/pkg/httpclient"
	"edge-common-go/pkg/log"
	"edge-common-go/pkg/utils"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

/**
 * @Author: qianjialin
 * @Date: 2020/8/11 17:00
 */
type NSQClient struct {
	Config      NSQClientConfig
	Consumer    *nsq.Consumer
	Producer    *nsq.Producer
	MsgChan     chan []byte
	Log         *log.LogAa
	srcFilename string
}

type NSQClientConfig struct {
	ClientId            string
	LookupdPollInterval int64
	Topic               string
	Channel             string
	Addr                []string
	ConnectMode         string //NSQD、NSQLOOKUPD、NSQDS
}

func (f *NSQClient) InitConsumer() error {
	f.srcFilename = utils.GetFileName()
	f.requestTopicCreateUntilSuccess()
	cfg := nsq.NewConfig()
	cfg.ClientID = f.Config.ClientId
	cfg.LookupdPollInterval = time.Duration(f.Config.LookupdPollInterval) * time.Millisecond
	var err error = nil
	f.Consumer, err = nsq.NewConsumer(f.Config.Topic, f.Config.Channel, cfg)
	if err != nil {
		f.Log.Errorf("[%s:%d]:connect nsq newConsumer error, addr=%s, clientId=%+v, err=%s", f.srcFilename, utils.GetLineNum(), f.Config.Addr, f.Config.ClientId, err)
		return err
	}
	f.Consumer.AddHandler(f)
	switch f.Config.ConnectMode {
	case "NSQD":
		//建立一个nsqd连接
		if err := f.Consumer.ConnectToNSQD(f.Config.Addr[0]); err != nil {
			f.Log.Errorf("[%s:%d]:connect nsq connectToNSQD error, addr=%s, clientId=%+v, err=%s", f.srcFilename, utils.GetLineNum(), f.Config.Addr, f.Config.ClientId, err)
			return err
		}
		break
	case "NSQLOOKUPD":
		//建立NSQLookupd连接
		if err := f.Consumer.ConnectToNSQLookupd(f.Config.Addr[0]); err != nil {
			f.Log.Errorf("[%s:%d]:connect nsq connectToNSQLookupd error, addr=%s, clientId=%+v, err=%s", f.srcFilename, utils.GetLineNum(), f.Config.Addr, f.Config.ClientId, err)
			return err
		}
		break
	case "NSQDS":
		//建立多个nsqd连接
		if err := f.Consumer.ConnectToNSQDs(f.Config.Addr); err != nil {
			f.Log.Errorf("[%s:%d]:connect nsq connectToNSQDs error, addr=%+v, clientId=%+v, err=%s", f.srcFilename, utils.GetLineNum(), f.Config.Addr, f.Config.ClientId, err)
			return err
		}
	}
	f.Log.Infof("[%s:%d]:connect nsq success, addr=%s, clientId=%+v", f.srcFilename, utils.GetLineNum(), f.Config.Addr, f.Config.ClientId)
	return nil
}

//处理消息
func (f *NSQClient) HandleMessage(msg *nsq.Message) error {
	//f.Log.Infof("client:%s receive from %s topic:%s, channel: %s, message: %s", f.Config.ClientId,
	//	msg.NSQDAddress, f.Config.Topic, f.Config.Channel, string(msg.Body))
	f.MsgChan <- msg.Body
	return nil
}

//请求NSQSubTopicCreateUrl直到成功为止
func (f *NSQClient) requestTopicCreateUntilSuccess() {
	//往每个nsqd发送创建请求
	for _, addr := range f.Config.Addr {
		var NSQSubTopicCreateUrl = fmt.Sprintf("http://%s/topic/create?topic=%s", addr, f.Config.Topic)
		var ticker *time.Ticker
		ticker = time.NewTicker(time.Second * 3)
		var err error
		for {
			select {
			case <-ticker.C:
				_, err = httpclient.Get(NSQSubTopicCreateUrl)
				if err == nil {
					ticker.Stop()
					return
				} else {
					f.Log.Errorf("[%s:%d]:connect nsq requestTopicCreateUrl error, url=%s,topic=%s, err=%s", f.srcFilename, utils.GetLineNum(), NSQSubTopicCreateUrl, f.Config.Topic, err)
				}
			}
		}
	}
}

/**
 ** 生产者*****************************************************************************************
 */
// 初始化生产者
func (f *NSQClient) InitProducer() error {
	f.srcFilename = utils.GetFileName()
	var err error
	cfg := nsq.NewConfig()
	cfg.ClientID = f.Config.ClientId
	f.Producer, err = nsq.NewProducer(f.Config.Addr[0], nsq.NewConfig())
	if err != nil {
		f.Log.Errorf("[%s:%d]:connect nsq newProducer error, addr=%s, clientId=%+v, err=%s", f.srcFilename, utils.GetLineNum(), f.Config.Addr, f.Config.ClientId, err)
		return err
	}
	f.Log.Infof("[%s:%d]:connect nsq success, addr=%s, clientId=%+v", f.srcFilename, utils.GetLineNum(), f.Config.Addr, f.Config.ClientId)
	return nil
}

//发布消息
func (f *NSQClient) Publish(topic string, message string) error {
	var err error
	if f.Producer != nil {
		if message == "" {
			f.Log.Warnf("[%s:%d]:nsq message = null", f.srcFilename, utils.GetLineNum())
			return nil
		}
		err = f.Producer.Publish(topic, []byte(message))
		if err != nil {
			f.Log.Errorf("[%s:%d]:nsq producer msg error, addr=%s, clientId=%+v, err=%s", f.srcFilename, utils.GetLineNum(), f.Config.Addr, f.Config.ClientId, err)
		}
		return err
	}
	f.Log.Errorf("[%s:%d]:nsq producer is nil error, addr=%s, clientId=%+v, err=%s", f.srcFilename, utils.GetLineNum(), f.Config.Addr, f.Config.ClientId, err)
	return fmt.Errorf("nsq producer is nil error,err=%s", err)
}
