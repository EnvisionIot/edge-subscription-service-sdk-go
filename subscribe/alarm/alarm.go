package alarm

import (
	"edge-subscription-go/utils"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

type AlarmHandler struct {
	logger  *log.Logger
	callBak func([]AlarmRecord)
}

func New(call func([]AlarmRecord)) *AlarmHandler {
	return &AlarmHandler{
		callBak: call,
		logger:  log.New(os.Stderr, "", log.Flags()),
	}
}

func (f *AlarmHandler) HandleMessage(msg *nsq.Message) error {
	var records AlarmRecords
	err := proto.Unmarshal(msg.Body, &records)
	if err != nil {
		msg := fmt.Sprintf("invalid AlarmRecords protobuf struct: %s", err.Error())
		_ = utils.CreateAndWriteSubLog(utils.ERROR, utils.GetFileNameLastN(1), utils.GetFuncName(), msg, utils.GetLineNum())
	}
	points := make([]AlarmRecord, len(records.Points))
	for index, record := range records.Points {
		points[index] = *record
	}
	f.callBak(points)

	return nil
}
