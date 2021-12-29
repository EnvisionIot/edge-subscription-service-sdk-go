/*
 * subscribe data from nsq, support realtimepoint, setmeasurepointresponse, controlresponse, alarm, event
 */

package subscribe

import (
	"edge-subscription-go/subscribe/alarm"
	"edge-subscription-go/subscribe/control"
	"edge-subscription-go/subscribe/record"
	"edge-subscription-go/subscribe/setpoint"
	"edge-subscription-go/utils"
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
)

type Subscribe struct {
	config   *utils.Subconfig
	consumer *nsq.Consumer
}

func init() {
	//_ = utils.NewGlobalLog()
}

func NewSubscriber(addr string, port int, topic, channel, accessKey, appSecret, subId string) (*Subscribe, error) {
	s := &Subscribe{}
	err := utils.NewGlobalLog()
	if err != nil {
		return nil, err
	}

	s.newConfig(addr, port, topic, channel, accessKey, appSecret, subId)
	err = s.newConsumer()
	return s, err
}

func (f *Subscribe) newConfig(addr string, port int, topic, channel, accessKey, appSecret, subId string) {
	f.config = &utils.Subconfig{
		IpAddr:      addr,
		Port:        port,
		TopicName:   topic,
		ChannelName: channel,
		AccessKey:   accessKey,
		AppSecret:   appSecret,
		SubId:       subId,
	}
}

func (f *Subscribe) newConsumer() error {
	msg, ok := edgeAuthorityCheck(f.config.AccessKey, f.config.AppSecret, f.config.SubId)
	if !ok {
		errStr := utils.CreateAndWriteSubLog(utils.ERROR, utils.GetFileNameLastN(1), utils.GetFuncName(), msg, utils.GetLineNum())
		return errors.New(errStr)
	}
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(f.config.TopicName, f.config.ChannelName, config)
	if err != nil {
		_ = utils.CreateAndWriteSubLog(utils.ERROR, utils.GetFileNameLastN(1), utils.GetFuncName(), err.Error(), utils.GetLineNum())
		return err
	}
	f.consumer = consumer

	return nil
}

func (f *Subscribe) connect() error {
	addr := fmt.Sprintf("%s:%d", f.config.IpAddr, f.config.Port)
	return f.consumer.ConnectToNSQD(addr)
}

//record point
func (f *Subscribe) SubRealTimePoint(call func([]record.SSRecordPoint)) error {
	handler := record.New(call)
	if f.consumer == nil {
		errStr := utils.CreateAndWriteSubLog(utils.ERROR, utils.GetFileNameLastN(1), utils.GetFuncName(), "invalid subscriber", utils.GetLineNum())
		return errors.New(errStr)
	}
	f.consumer.AddHandler(handler)

	return f.connect()
}

//alarm data
func (f *Subscribe) SubAlarmData(call func([]alarm.AlarmRecord)) error {
	handler := alarm.New(call)
	if f.consumer == nil {
		errStr := utils.CreateAndWriteSubLog(utils.ERROR, utils.GetFileNameLastN(1), utils.GetFuncName(), "invalid subscriber", utils.GetLineNum())
		return errors.New(errStr)
	}
	f.consumer.AddHandler(handler)

	return f.connect()
}

//control response
func (f *Subscribe) SubControlResponse(call func([]control.ControlResponsePoint)) error {
	handler := control.New(call)
	if f.consumer == nil {
		errStr := utils.CreateAndWriteSubLog(utils.ERROR, utils.GetFileNameLastN(1), utils.GetFuncName(), "invalid subscriber", utils.GetLineNum())
		return errors.New(errStr)
	}
	f.consumer.AddHandler(handler)
	return f.connect()
}

//set point response
func (f *Subscribe) SubSetMeasurePointResponse(call func([]setpoint.SetMeasurepointResponsePoint)) error {
	handler := setpoint.New(call)
	if f.consumer == nil {
		errStr := utils.CreateAndWriteSubLog(utils.ERROR, utils.GetFileNameLastN(1), utils.GetFuncName(), "invalid subscriber", utils.GetLineNum())
		return errors.New(errStr)
	}
	f.consumer.AddHandler(handler)
	return f.connect()
}
