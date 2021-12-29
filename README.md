# Using EnOS Edge Data Subscription SDK for Go

EnOS Edge Data Subscription Service improves the API calling efficiency of applications with active data push, which supports subscription to real-time measurment point, the response of setting measurement point and control operations, and alert data.

After configuring and starting data subscription jobs on the EnOS Management Console, you can use the Data Subscription SDK for Go to develop applications for consuming the subscribed data.

## Installation

### Prerequisites

The Data Subscription SDK for Go supports Go 1.12.5.

### Building From Source

Obtain the source code of Data Subscription SDK for go.

- From GitHub:

```
git clone https://github.com/EnvisionIot/edge-subscription-service-sdk-go
```
- From EnOS SDK Center. Click **SDK Center** from the right navigation of EnOS Console, and obtain the SDK source code by clicking the GitHub icon in the **Obtain** column.

You can install the SDK from source.

### Initialize Subscriber with Function
```
import (
	"edge-subscription-go/subscribe"
)

func main(){
	edgeServiceIp := "127.0.0.1"
	edgeServicePort := 9150
	subTopicName := "DATASVC.SUB.group"
	subChannelName := "real_time_point"
	accessKey := "a-b-c"
	appSecret := "a-d-f"
	subId := "subTopicId"
	f, err := subscribe.NewSubscriber(edgeServiceIp, edgeServicePort, subTopicName, subChannelName,accessKey, appSecret, subId)
	if err != nil{
		//init failed, return
		return
	}
}
```

### Define the Message Processing Function
```
import (
	"edge-subscription-go/subscribe/record"
	"fmt"
)

func main(){
    ...
	functionRealTimePoint := func(records []record.SSRecordPoint) {
    		for _, point := range records{
    			fmt.Printf("real time point value = %s\n", point.GetValue())
    		}
    }
    ...
}
```

## Feature List

EnOS Enterprise Data Platform supports subscribing to asset time series data and alert data and pushing the subscribed data to applications, thus improving the data query efficiency of applications.

The features supported by this SDK include:

- Consuming subscribed real-time data
- Consuming subscribed alert data
- Consuming subscribed response of setting points
- Consuming subscribed results of control command

## Sample Code

### Code Sample for Consuming Subscribed Real-time Data
```
import (
	"edge-subscription-go/subscribe"
	"edge-subscription-go/subscribe/record"
	"fmt"
)

func main(){
	//Initialize Subscriber
	edgeServiceIp := "127.0.0.1"
	edgeServicePort := 9150
	subTopicName := "DATASVC.SUB.group"
	subChannelName := "real_time_point"
	accessKey := "accessKey"
	appSecret := "appSecret"
	subId := "subTopicId"
	f, err := subscribe.NewSubscriber(edgeServiceIp, edgeServicePort, subTopicName, subChannelName,accessKey, appSecret, subId)
	if err != nil{
		return
	}

	//Define the Message Processing Function
	functionRealTimePoint := func(records []record.SSRecordPoint) {
		for _, point := range records{
			fmt.Printf("real time point value = %s\n", point.GetValue())
		}
	}
	//start
	_ = f.SubRealTimePoint(functionRealTimePoint)
}
```

### Code Sample for Consuming Subscribed Alert Data
```
import (
	"edge-subscription-go/subscribe"
	"edge-subscription-go/subscribe/alarm"
	"fmt"
)

func main(){
	//Initialize Subscriber
	edgeServiceIp := "127.0.0.1"
	edgeServicePort := 9150
	subTopicName := "DATASVC.SUB.group"
	subChannelName := "alarm_data"
	accessKey := "accessKey"
	appSecret := "appSecret"
	subId := "subTopicId"
	f, err := subscribe.NewSubscriber(edgeServiceIp, edgeServicePort, subTopicName, subChannelName,accessKey, appSecret, subId)
	if err != nil{
		return
	}

	//Define the Message Processing Function
	functionAlarmData := func(records []alarm.AlarmRecord) {
		for _, point := range records{
			fmt.Printf("alarm data value = %s\n", point.GetValue())
		}
	}
	//start
	_ = f.SubAlarmData(functionAlarmData)
}
```

### Code Sample for Consuming Subscribed Response of Setting Points
```
import (
	"edge-subscription-go/subscribe"
	"edge-subscription-go/subscribe/control"
	"fmt"
)

func main(){
	//Initialize Subscriber
	edgeServiceIp := "127.0.0.1"
	edgeServicePort := 9150
	subTopicName := "DATASVC.SUB.group"
	subChannelName := "alarm_data"
	accessKey := "accessKey"
	appSecret := "appSecret"
	subId := "subTopicId"
	f, err := subscribe.NewSubscriber(edgeServiceIp, edgeServicePort, subTopicName, subChannelName,accessKey, appSecret, subId)
	if err != nil{
		return
	}

	//Define the Message Processing Function
	functionSetPointResponse := func(records []setpoint.SetMeasurepointResponsePoint) {
		for _, point := range records{
			fmt.Printf("set measurePoint input value = %s\n", point.GetInputData())
		}
	}
	//start
	_ = f.SubSetMeasurePointResponse(functionSetPointResponse)
}
```

### Code Sample for Consuming Subscribed Results of Control Command
```
import (
	"edge-subscription-go/subscribe"
	"edge-subscription-go/subscribe/control"
	"fmt"
)

func main(){
	//Initialize Subscriber
	edgeServiceIp := "127.0.0.1"
	edgeServicePort := 9150
	subTopicName := "DATASVC.SUB.group"
	subChannelName := "alarm_data"
	accessKey := "accessKey"
	appSecret := "appSecret"
	subId := "subTopicId"
	f, err := subscribe.NewSubscriber(edgeServiceIp, edgeServicePort, subTopicName, subChannelName,accessKey, appSecret, subId)
	if err != nil{
		return
	}

	//Define the Message Processing Function
	functionControlResponse := func(records []control.ControlResponsePoint) {
		for _, point := range records{
			fmt.Printf("control response value = %s\n", point.GetInputData())
		}
	}
	//start
	_ = f.SubControlResponse(functionControlResponse)
}
```
