package utils

const (
	IamAddr     = "http://edgeextensive001:8092/apim-token-service/v2.0/token/checkSub"
	HttpTimeOut = 2000
	LogFormat   = "[%s:%s:%d]:%v" //file:function:line
)

type LogDisk int

const (
	DEBUG = LogDisk(1)
	INFO  = LogDisk(2)
	WARN  = LogDisk(3)
	ERROR = LogDisk(4)
	FATAL = LogDisk(5)
)

type Subconfig struct {
	IpAddr      string
	Port        int
	TopicName   string
	ChannelName string
	AccessKey   string
	AppSecret   string
	SubId       string
}

type IamParam struct {
	AppKey     string `json:"appKey"`
	Encryption string `json:"encryption"`
	Timestamp  int64  `json:"timestamp"`
	SubId      string `json:"subId"`
	SkipSubId  bool   `json:"skipSubId"`
}

type IamRsp struct {
	Status   int         `json:"status"`
	Msg      string      `json:"msg"`
	Business string      `json:"business"`
	Data     interface{} `json:"data"`
}
