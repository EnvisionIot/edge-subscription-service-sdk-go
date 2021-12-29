package utils

import (
	"edge-common-go/pkg/log"
	"edge-common-go/pkg/utils"
	"fmt"
)

var (
	GlobalLog *log.LogAa
)

func NewGlobalLog() error {
	err := log.LogAGlobal.Init(utils.GetBaseDir())
	if err != nil {
		return err
	}
	GlobalLog, err = log.GetLogAa("/sub.log")
	return err
}

func CreateAndWriteSubLog(disk LogDisk, fileName, functionName, msg string, line int64) string {
	str := fmt.Sprintf(LogFormat, fileName, functionName, line, msg)
	switch disk {
	case INFO:
		GlobalLog.Info(str)
	case ERROR:
		GlobalLog.Error(str)
	case WARN:
		GlobalLog.Warn(str)
	case DEBUG:
		GlobalLog.Debug(str)
	case FATAL:
		GlobalLog.Fatal(str)
	default:
		GlobalLog.Info(str)
	}

	return str
}
