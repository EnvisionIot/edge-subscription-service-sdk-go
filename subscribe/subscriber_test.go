package subscribe

import (
	"edge-subscription-go/subscribe/alarm"
	"edge-subscription-go/subscribe/control"
	"edge-subscription-go/subscribe/record"
	"edge-subscription-go/subscribe/setpoint"
	"fmt"
	"testing"
	"time"
)

func TestNewSubscriber(t *testing.T) {
	_, err := NewSubscriber("127.0.0.1", 4150, "subscriber", "real_time_point", "", "", "")
	if err != nil {
		t.FailNow()
	}
}

func TestSubscribe_SubRealTimePoint(t *testing.T) {
	f, err := NewSubscriber("127.0.0.1", 4150, "DATASVC.SUB.grouptest11111", "real_time_point_jack", "0ae4c1ab-605c-4e36-b2b3-aab38026ac57", "520351be-2347-4399-9831-fbafb5ca11af", "DATASVC.SUB.grouptest11111")
	if err != nil {
		fmt.Printf("NewSubscriber error: %s", err.Error())
		return
	}
	function := func(records []record.SSRecordPoint) {
		for _, point := range records {
			fmt.Printf("record value = %s\n", point.GetValue())
		}
	}
	_ = f.SubRealTimePoint(function)
	time.Sleep(time.Second * 2)
}

func TestSubscribe_SubAlarmData(t *testing.T) {
	f, err := NewSubscriber("127.0.0.1", 4150, "subscriber_alarm", "alarm_data", "", "", "")
	if err != nil {
		fmt.Printf("NewSubscriber error: %s", err.Error())
		return
	}
	function := func(records []alarm.AlarmRecord) {
		for _, point := range records {
			fmt.Printf("record value = %s\n", point.GetValue())
		}
	}
	_ = f.SubAlarmData(function)
	time.Sleep(time.Second * 2)
}

func TestSubscribe_SubControlResponse(t *testing.T) {
	f, err := NewSubscriber("127.0.0.1", 4150, "subscriber_control", "control_response", "", "", "")
	if err != nil {
		fmt.Printf("NewSubscriber error: %s", err.Error())
		return
	}
	function := func(records []control.ControlResponsePoint) {
		for _, point := range records {
			fmt.Printf("record value = %s\n", point.GetInputData())
		}
	}
	_ = f.SubControlResponse(function)
	time.Sleep(time.Second * 2)
}

func TestSubscribe_SubSetMeasurePointResponse(t *testing.T) {
	f, err := NewSubscriber("127.0.0.1", 4150, "subscriber_set_point", "set_point_response", "", "", "")
	if err != nil {
		fmt.Printf("NewSubscriber error: %s", err.Error())
		return
	}
	function := func(records []setpoint.SetMeasurepointResponsePoint) {
		for _, point := range records {
			fmt.Printf("record input value = %s\n", point.GetInputData())
		}
	}
	_ = f.SubSetMeasurePointResponse(function)
	time.Sleep(time.Second * 2)
}
