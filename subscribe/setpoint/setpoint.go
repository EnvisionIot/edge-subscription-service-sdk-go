package setpoint

import (
	"edge-subscription-go/utils"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

type SetPointHandler struct {
	callBak func([]SetMeasurepointResponsePoint)
	logger  *log.Logger
}

func New(call func([]SetMeasurepointResponsePoint)) *SetPointHandler {
	return &SetPointHandler{
		logger:  log.New(os.Stderr, "", log.Flags()),
		callBak: call,
	}
}

func (f *SetPointHandler) HandleMessage(msg *nsq.Message) error {
	var records SetMeasurepointResponsePoints
	err := proto.Unmarshal(msg.Body, &records)
	if err != nil {
		msg := fmt.Sprintf("invalid SetMeasurepointResponsePoints protobuf struct: %s", err.Error())
		_ = utils.CreateAndWriteSubLog(utils.ERROR, utils.GetFileNameLastN(1), utils.GetFuncName(), msg, utils.GetLineNum())
	}
	points := make([]SetMeasurepointResponsePoint, len(records.Points))
	for index, record := range records.Points {
		points[index] = *record
	}
	f.callBak(points)

	return nil
}
