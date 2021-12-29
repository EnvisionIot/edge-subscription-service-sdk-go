package record

import (
	"fmt"
	"github.com/edge-subscription-service-sdk-go/utils"
	"github.com/golang/protobuf/proto"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

type PointHandler struct {
	callBak func([]SSRecordPoint)
	logger  *log.Logger
}

func New(call func([]SSRecordPoint)) *PointHandler {
	return &PointHandler{
		logger:  log.New(os.Stderr, "", log.Flags()),
		callBak: call,
	}
}

func (f *PointHandler) HandleMessage(msg *nsq.Message) error {
	var records SSRecordPoints
	err := proto.Unmarshal(msg.Body, &records)
	if err != nil {
		msg := fmt.Sprintf("invalid SSRecordPoints protobuf struct: %s", err.Error())
		_ = utils.CreateAndWriteSubLog(utils.ERROR, utils.GetFileNameLastN(1), utils.GetFuncName(), msg, utils.GetLineNum())
	}
	points := make([]SSRecordPoint, len(records.Points))
	for index, record := range records.Points {
		points[index] = *record
	}
	f.callBak(points)

	return nil
}
