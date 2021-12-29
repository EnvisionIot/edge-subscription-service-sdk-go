package main

import (
	"edge-subscription-go/subscribe"
	"edge-subscription-go/subscribe/alarm"
	"edge-subscription-go/subscribe/control"
	"edge-subscription-go/subscribe/record"
	"edge-subscription-go/subscribe/setpoint"
	"fmt"
)

func main() {
	//1.Initialize a new subscriber with function: NewSubscriber.
	edgeServiceIp := "127.0.0.1"
	edgeServicePort := 9150
	subTopicName := "DATASVC.SUB.group"
	subChannelName := "real_time_point"
	accessKey := "a-b-c"
	appSecret := "a-d-f"
	subId := "subTopicId"
	f, err := subscribe.NewSubscriber(edgeServiceIp, edgeServicePort, subTopicName, subChannelName, accessKey, appSecret, subId)
	if err != nil {
		//init failed, return
		return
	}

	//Subscribe Real-time measuring point from Edge service
	//2.Define the message processing function according to the subscription type.
	functionRealTimePoint := func(records []record.SSRecordPoint) {
		for _, point := range records {
			fmt.Printf("real time point value = %s\n", point.GetValue())
		}
	}
	//3.Start subscribing, receive data.
	_ = f.SubRealTimePoint(functionRealTimePoint)

	//Subscribe alarm data from Edge service
	functionAlarmData := func(records []alarm.AlarmRecord) {
		for _, point := range records {
			fmt.Printf("alarm data value = %s\n", point.GetValue())
		}
	}
	_ = f.SubAlarmData(functionAlarmData)

	//Subscribe control response data from Edge service
	functionControlResponse := func(records []control.ControlResponsePoint) {
		for _, point := range records {
			fmt.Printf("control response value = %s\n", point.GetInputData())
		}
	}
	_ = f.SubControlResponse(functionControlResponse)

	//Subscribe SetMeasurePointResponse data from Edge service
	functionSetPointResponse := func(records []setpoint.SetMeasurepointResponsePoint) {
		for _, point := range records {
			fmt.Printf("set measurePoint input value = %s\n", point.GetInputData())
		}
	}
	_ = f.SubSetMeasurePointResponse(functionSetPointResponse)
}
