package log

/**
* Description:
*
* @author yang.zhang4
* @date 2019/6/15
 */

import (
	rotatelogs "edge-common-go/pkg/log/lestrrat-go/file-rotatelogs"
	"encoding/json"
	"fmt"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	DefaultLogLevel         = "info"
	DefaultLogDir           = "logs"
	DefaultDefaultLogPath   = "edge-defalut.log"
	DefaultMonitorLogPath   = "_default.log"
	DefaultRotationTime     = 86400
	DefaultMaxAge           = 604800
	DefaultRotationCount    = -1
	DefaultMaxFileKeepCount = 10
	DefaultMaxFileSize      = 100

	DefaultLogConfigPath = "resources/log.conf"

	PanicLevelStr = "panic"
	FatalLevelStr = "fatal"
	ErrorLevelStr = "error"
	WarnLevelStr  = "warn"
	InfoLevelStr  = "info"
	DebugLevelStr = "debug"

	PanicLevel = 1
	FatalLevel = 2
	ErrorLevel = 3
	WarnLevel  = 4
	InfoLevel  = 5
	DebugLevel = 6
)

var LogAGlobal LogA

type LogA struct {
	LogAaMap     map[string]*LogAa
	Config       LogAConfig
	ConfigFormat LogAConfigFormat
	BaseDir      string
	//操作LogAaMap的锁
	aaMutex sync.Mutex
	//操作ConfigFormat的锁
	bbMutex sync.Mutex
}

type LogAa struct {
	filePath string
	logger   *logrus.Logger
	//操作Level的锁
	ccMutex sync.RWMutex
}

type LogAConfig struct {
	LogLevel              string
	LogDir                string
	DefaultLogPath        string
	DefaultMonitorLogPath string
	RotationTime          int64
	MaxAge                int64
	RotationCount         int64
	MaxFileKeepCount      int
	MaxFileSize           int
}

type LogAConfigFormat struct {
	LogLevel              string
	LogLevelFormat        int32
	LogDir                string
	DefaultLogPath        string
	DefaultMonitorLogPath string
	RotationTime          int64
	MaxAge                int64
	RotationCount         int64
	MaxFileKeepCount      int
	MaxFileSize           int64
}

func (l *LogA) Init(baseDir string) error {
	l.LogAaMap = make(map[string]*LogAa)
	l.BaseDir = baseDir
	l.Config.LogLevel = DefaultLogLevel
	l.Config.LogDir = DefaultLogDir
	l.Config.DefaultLogPath = DefaultDefaultLogPath
	l.Config.DefaultMonitorLogPath = DefaultMonitorLogPath
	l.Config.RotationTime = DefaultRotationTime
	l.Config.MaxAge = DefaultMaxAge
	l.Config.RotationCount = DefaultRotationCount
	l.Config.MaxFileKeepCount = DefaultMaxFileKeepCount
	l.Config.MaxFileSize = DefaultMaxFileSize

	var err error
	l.bbMutex.Lock()
	defer l.bbMutex.Unlock()
	l.ConfigFormat, err = logAConfigToLogAConfigFormat(l.Config)
	if err != nil {
		return err
	}

	logConfigFilePath := baseDir + "/" + DefaultLogConfigPath
	data, err := ioutil.ReadFile(logConfigFilePath)
	if err == nil {
		err = json.Unmarshal(data, &(l.Config))
		if err != nil {
			return err
		}
	}
	
	if l.Config.MaxFileKeepCount == 0 {
		l.Config.MaxFileKeepCount = DefaultMaxFileKeepCount
	}
	if l.Config.MaxFileSize == 0 {
		l.Config.MaxFileSize = DefaultMaxFileSize
	}

	l.ConfigFormat, err = logAConfigToLogAConfigFormat(l.Config)
	if err != nil {
		return err
	}

	go l.logConfigJob()

	return nil
}

func GetLogAa(filePath ...string) (*LogAa, error) {
	return LogAGlobal.GetLogAa(filePath...)
}

func (l *LogA) GetLogAa(filePath ...string) (*LogAa, error) {
	var logFilePath string
	l.bbMutex.Lock()
	defer l.bbMutex.Unlock()
	if len(filePath) == 0 {
		logFilePath = l.BaseDir + "/" + l.ConfigFormat.LogDir + "/" + l.ConfigFormat.DefaultLogPath
	} else {
		str := strings.Trim(filePath[0], " \r\n")
		if len(str) <= 0 {
			logFilePath = l.BaseDir + "/" + l.ConfigFormat.LogDir + "/" + l.ConfigFormat.DefaultLogPath
		} else {
			logFilePath = l.BaseDir + "/" + l.ConfigFormat.LogDir + "/" + filePath[0]
		}
	}

	var logDir string
	logDir = filepath.Dir(logFilePath)
	isExist, _ := PathExists(logDir)
	if !isExist {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			return nil, fmt.Errorf("can not create dir %v", logDir)
		}
	}

	l.aaMutex.Lock()
	defer l.aaMutex.Unlock()

	v, ok := l.LogAaMap[logFilePath]
	if ok {
		return v, nil
	}
	aa, err := l.creatLogAa(logFilePath)
	if err != nil {
		return aa, err
	}
	l.LogAaMap[logFilePath] = aa

	return aa, nil
}

//获取记录监控日志的logAa，里面的日志用于falcon监控等
func GetLogAaMonitor() (*LogAa, error) {
	return GetLogAa(LogAGlobal.ConfigFormat.DefaultMonitorLogPath)
}

//将日志记录进监控文件
func LogToMonitorErrorf(format string, args ...interface{}) {
	var l *LogAa
	l, _ = GetLogAaMonitor()
	l.logger.Errorf(format, args...)
}

//将日志记录进监控文件
func LogToMonitorError(args ...interface{}) {
	var l *LogAa
	l, _ = GetLogAaMonitor()
	l.logger.Error(args...)
}

//将日志记录到多个log-->
func MultiLogErrorf(logArray []*LogAa, format string, args ...interface{}) {
	for _, v := range logArray {
		v.Errorf(format, args...)
	}
}

func MultiLogError(logArray []*LogAa, args ...interface{}) {
	for _, v := range logArray {
		v.Error(args...)
	}
}

func MultiLogInfof(logArray []*LogAa, format string, args ...interface{}) {
	for _, v := range logArray {
		v.Infof(format, args...)
	}
}

func MultiLogInfo(logArray []*LogAa, args ...interface{}) {
	for _, v := range logArray {
		v.Info(args...)
	}
}

func MultiLogWarnf(logArray []*LogAa, format string, args ...interface{}) {
	for _, v := range logArray {
		v.Warnf(format, args...)
	}
}

func MultiLogWarn(logArray []*LogAa, args ...interface{}) {
	for _, v := range logArray {
		v.Warn(args...)
	}
}

func MultiLogDebugf(logArray []*LogAa, format string, args ...interface{}) {
	for _, v := range logArray {
		v.Debugf(format, args...)
	}
}

func MultiLogDebug(logArray []*LogAa, args ...interface{}) {
	for _, v := range logArray {
		v.Debug(args...)
	}
}

//<--将日志记录到多个log

func (aa *LogAa) Panicf(format string, args ...interface{}) {
	aa.logger.Panicf(format, args...)
}

func (aa *LogAa) Panic(args ...interface{}) {
	aa.logger.Panic(args...)
}

func (aa *LogAa) Panicln(args ...interface{}) {
	aa.logger.Panicln(args...)
}

func (aa *LogAa) Fatalf(format string, args ...interface{}) {
	aa.logger.Fatalf(format, args...)
}

func (aa *LogAa) Fatal(args ...interface{}) {
	aa.logger.Fatal(args...)
}

func (aa *LogAa) Fatalln(args ...interface{}) {
	aa.logger.Fatalln(args...)
}

func (aa *LogAa) Errorf(format string, args ...interface{}) {
	aa.logger.Errorf(format, args...)
}

func (aa *LogAa) Error(args ...interface{}) {
	aa.logger.Error(args...)
}

func (aa *LogAa) Errorln(args ...interface{}) {
	aa.logger.Errorln(args...)
}

func (aa *LogAa) Warnf(format string, args ...interface{}) {
	aa.logger.Warnf(format, args...)
}

func (aa *LogAa) Warn(args ...interface{}) {
	aa.logger.Warn(args...)
}

func (aa *LogAa) Warnln(args ...interface{}) {
	aa.logger.Warnln(args...)
}

func (aa *LogAa) Infof(format string, args ...interface{}) {
	aa.logger.Infof(format, args...)
}

func (aa *LogAa) Info(args ...interface{}) {
	aa.logger.Info(args...)
}

func (aa *LogAa) Infoln(args ...interface{}) {
	aa.logger.Infoln(args...)
}

func (aa *LogAa) Debugf(format string, args ...interface{}) {
	aa.logger.Debugf(format, args...)
}

func (aa *LogAa) Debug(args ...interface{}) {
	aa.logger.Debug(args...)
}

func (aa *LogAa) Debugln(args ...interface{}) {
	aa.logger.Debugln(args...)
}

func (aa *LogAa) IsDebugEnabled() bool {
	aa.ccMutex.RLock()
	defer aa.ccMutex.RUnlock()
	if int32(aa.logger.Level) >= int32(logrus.DebugLevel) {
		return true
	} else {
		return false
	}
}

func (aa *LogAa) IsInfoEnabled() bool {
	aa.ccMutex.RLock()
	defer aa.ccMutex.RUnlock()
	if int32(aa.logger.Level) >= int32(logrus.InfoLevel) {
		return true
	} else {
		return false
	}
}

func (aa *LogAa) IsWarnEnabled() bool {
	aa.ccMutex.RLock()
	defer aa.ccMutex.RUnlock()
	if int32(aa.logger.Level) >= int32(logrus.WarnLevel) {
		return true
	} else {
		return false
	}
}

func (aa *LogAa) IsErrorEnabled() bool {
	aa.ccMutex.RLock()
	defer aa.ccMutex.RUnlock()
	if int32(aa.logger.Level) >= int32(logrus.ErrorLevel) {
		return true
	} else {
		return false
	}
}

func (aa *LogAa) IsFatalEnabled() bool {
	aa.ccMutex.RLock()
	defer aa.ccMutex.RUnlock()
	if int32(aa.logger.Level) >= int32(logrus.FatalLevel) {
		return true
	} else {
		return false
	}
}

func (aa *LogAa) IsPanicEnabled() bool {
	aa.ccMutex.RLock()
	defer aa.ccMutex.RUnlock()
	if int32(aa.logger.Level) >= int32(logrus.PanicLevel) {
		return true
	} else {
		return false
	}
}

func (l *LogA) creatLogAa(filePath string) (*LogAa, error) {
	a := LogAa{}
	a.filePath = filePath
	a.logger = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(EdgeTextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     InfoLevel,
	}

	var rotationCountOption rotatelogs.Option
	if l.ConfigFormat.RotationCount <= 0 {
		rotationCountOption = rotatelogs.WithRotationCount(0)
	} else {
		rotationCountOption = rotatelogs.WithRotationCount(uint(l.ConfigFormat.RotationCount))
	}

	writer, err := rotatelogs.New(
		filePath+".%Y%m%d%H%M",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(filePath),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Duration(l.ConfigFormat.RotationTime)),

		// WithMaxAge和WithRotationCount二者只能设置一个，设为-1表示不生效
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(time.Duration(l.ConfigFormat.MaxAge)),
		rotationCountOption,
		rotatelogs.WithClock(rotatelogs.Local),

		//file size
		rotatelogs.WithMaxFileKeepCount(l.ConfigFormat.MaxFileKeepCount),
		rotatelogs.WithMaxFileSize(l.ConfigFormat.MaxFileSize),
	)

	if err != nil {
		return &a, fmt.Errorf("config local file system for logger error: %v", err)
	}

	lvl, _ := logrus.ParseLevel(l.ConfigFormat.LogLevel)
	a.logger.SetLevel(lvl)

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &EdgeTextFormatter{})

	a.logger.Hooks.Add(lfsHook)

	return &a, nil
}

type EdgeTextFormatter struct {
}

func (f *EdgeTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("[%s][%s] - %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func logLevelStrToInt32(level string) (int32, error) {
	var ret int32 = InfoLevel
	switch level {
	case DebugLevelStr:
		ret = DebugLevel
	case InfoLevelStr:
		ret = InfoLevel
	case WarnLevelStr:
		ret = WarnLevel
	case ErrorLevelStr:
		ret = ErrorLevel
	case FatalLevelStr:
		ret = FatalLevel
	case PanicLevelStr:
		ret = PanicLevel
	default:
		ret = InfoLevel
		return ret, fmt.Errorf("unsupported LogLevel")
	}

	return ret, nil
}

func logAConfigToLogAConfigFormat(l LogAConfig) (LogAConfigFormat, error) {
	lf := LogAConfigFormat{}
	var str string

	str = strings.Trim(l.LogLevel, " \r\n")
	if len(str) <= 0 {
		lf.LogLevel = DefaultLogLevel
	} else {
		lf.LogLevel = str
	}
	lf.LogLevelFormat, _ = logLevelStrToInt32(lf.LogLevel)

	str = strings.Trim(l.LogDir, " \r\n")
	if len(str) <= 0 {
		lf.LogDir = DefaultLogDir
	} else {
		lf.LogDir = str
	}

	str = strings.Trim(l.DefaultLogPath, " \r\n")
	if len(str) <= 0 {
		lf.DefaultLogPath = DefaultDefaultLogPath
	} else {
		lf.DefaultLogPath = str
	}

	str = strings.Trim(l.DefaultMonitorLogPath, " \r\n")
	if len(str) <= 0 {
		lf.DefaultMonitorLogPath = DefaultMonitorLogPath
	} else {
		lf.DefaultMonitorLogPath = str
	}

	if l.RotationTime <= 0 {
		lf.RotationTime = DefaultRotationTime * int64(time.Second)
	} else {
		lf.RotationTime = l.RotationTime * int64(time.Second)
	}

	if l.MaxAge <= 0 {
		lf.MaxAge = -1
	} else {
		lf.MaxAge = l.MaxAge * int64(time.Second)
	}

	if l.RotationCount <= 0 {
		lf.RotationCount = -1
	} else {
		lf.RotationCount = l.RotationCount
	}

	if lf.MaxAge < 0 && lf.RotationCount < 0 {
		lf.MaxAge = DefaultMaxAge * int64(time.Second)
	}
	if l.MaxFileSize >= 200 {
		lf.MaxFileSize = 200 * 1024 * 1024
	} else if l.MaxFileSize <= 0 {
		lf.MaxFileSize = 100 * 1024 * 1024
	} else {
		lf.MaxFileSize = int64(l.MaxFileSize) * 1024 * 1024
	}
	if l.MaxFileKeepCount >= 20 {
		lf.MaxFileKeepCount = 20
	} else if l.MaxFileKeepCount <= 0 {
		lf.MaxFileKeepCount = 10
	} else {
		lf.MaxFileKeepCount = l.MaxFileKeepCount
	}

	return lf, nil
}

//判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//周期扫描配置文件
func (l *LogA) logConfigJob() {
	var timerTime time.Duration = 10 * time.Second
	timer := time.NewTicker(timerTime)

	for {
		select {
		case <-timer.C:
			l.parseConfigFile()
		}
	}
}

func (l *LogA) parseConfigFile() {
	var err error
	var data []byte
	var logConfig LogAConfig

	logConfigFilePath := l.BaseDir + "/" + DefaultLogConfigPath
	data, err = ioutil.ReadFile(logConfigFilePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &(logConfig))
	if err != nil {
		return
	}

	err = l.parseConfig(&logConfig)
	if err != nil {
		return
	}
}

func (l *LogA) parseConfig(logConfig *LogAConfig) error {
	if logConfig == nil {
		return fmt.Errorf("logConfig is nil error")
	}
	var logLevel = DefaultLogLevel
	str := strings.Trim((*logConfig).LogLevel, " \r\n")
	if len(str) <= 0 {
		logLevel = DefaultLogLevel
	} else {
		logLevel = str
	}
	if l.ConfigFormat.LogLevel == logLevel {
		//没变无需设置level
		return nil
	}
	l.bbMutex.Lock()
	defer l.bbMutex.Unlock()
	l.ConfigFormat.LogLevel = logLevel

	l.aaMutex.Lock()
	defer l.aaMutex.Unlock()
	lvl, _ := logrus.ParseLevel(logLevel)
	for _, log := range l.LogAaMap {
		log.ccMutex.Lock()
		log.logger.SetLevel(lvl)
		log.ccMutex.Unlock()
	}
	return nil
}
