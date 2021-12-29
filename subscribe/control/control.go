package control

import (
	"edge-subscription-go/utils"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

type ControlHandler struct {
	callBak func([]ControlResponsePoint)
	logger  *log.Logger
}

func New(call func([]ControlResponsePoint)) *ControlHandler {
	return &ControlHandler{
		logger:  log.New(os.Stderr, "", log.Flags()),
		callBak: call,
	}
}

func (f *ControlHandler) HandleMessage(msg *nsq.Message) error {
	var records ControlResponsePoints
	err := proto.Unmarshal(msg.Body, &records)
	if err != nil {
		msg := fmt.Sprintf("invalid ControlResponsePoint protobuf struct: %s", err.Error())
		_ = utils.CreateAndWriteSubLog(utils.ERROR, utils.GetFileNameLastN(1), utils.GetFuncName(), msg, utils.GetLineNum())
	}
	points := make([]ControlResponsePoint, len(records.Points))
	for index, record := range records.Points {
		points[index] = *record
	}
	f.callBak(points)

	return nil
}
